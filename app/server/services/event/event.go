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
	ChannelName = "p2p-peer-registry-cmd"
)

const KubemqDefaultPort = 50000

type EventEmit interface {
	EmitP2pPeerRegistryCmd(queueUrl string, clientId string, event *model.P2pPeerRegistryCmd) (string, error)
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

func (*EventEmitImpl) EmitP2pPeerRegistryCmd(queueUrl string, clientId string, event *model.P2pPeerRegistryCmd) (string, error) {
	log.Infoln("EmitP2pPeerRegistryCmd")
	bEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatalln(err)
	}
	msgId, err := emitEvent(queueUrl, clientId, bEvent, ChannelName)
	if err != nil {
		log.Fatalln(err)
	}
	return msgId, nil
}
