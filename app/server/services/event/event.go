package event

import (
	"context"
	"encoding/json"

	"blitzshare.api/app/model"
	"github.com/kubemq-io/kubemq-go"
	log "github.com/sirupsen/logrus"
)

type NodeJoinedEvent struct {
	NodeId string `json:"nodeId"`
}

func NewNodeJoinedEvent(nodeId string) *NodeJoinedEvent {
	return &NodeJoinedEvent{NodeId: nodeId}
}

const (
	PeerRegisterCmd   = "p2p-peer-register-cmd"
	PeerDeregisterCmd = "p2p-peer-deregister-cmd"
)

const KubemqDefaultPort = 50000

type EventEmit interface {
	EmitP2pPeerRegisterCmd(queueUrl string, clientId string, event *model.P2pPeerRegistryCmd) (string, error)
	EmitP2pPeerDeregisterCmd(queueUrl string, clientId string, event *model.P2pPeerDeregisterCmd) (string, error)
}

type EventEmitImpl struct {
}

func NewEventEmit() EventEmit {
	return &EventEmitImpl{}
}

func emitEvent(queueUrl string, clientId string, event []byte, channelName string) (string, error) {
	ctx, _ := context.WithCancel(context.Background())
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(queueUrl, KubemqDefaultPort),
		kubemq.WithClientId(clientId),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	defer client.Close()
	if err != nil {
		log.Fatalln("cant connect to queue", queueUrl, err)
	}
	log.Infoln("emitEvent", channelName)
	sendResult, err := client.NewQueueMessage().
		SetChannel(channelName).
		SetBody(event).
		Send(ctx)
	return sendResult.MessageID, err
}

func (*EventEmitImpl) EmitP2pPeerRegisterCmd(queueUrl string, clientId string, event *model.P2pPeerRegistryCmd) (string, error) {
	log.Debugln("EmitP2pPeerRegisterCmd")
	bEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatalln(err)
	}
	msgId, err := emitEvent(queueUrl, clientId, bEvent, PeerRegisterCmd)
	if err != nil {
		log.Fatalln(err)
	}
	return msgId, nil
}

func (*EventEmitImpl) EmitP2pPeerDeregisterCmd(queueUrl string, clientId string, event *model.P2pPeerDeregisterCmd) (string, error) {
	log.Debugln("EmitP2pPeerDeregisterCmd")
	bEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatalln(err)
	}
	msgId, err := emitEvent(queueUrl, clientId, bEvent, PeerDeregisterCmd)
	if err != nil {
		log.Fatalln(err)
	}
	return msgId, nil
}
