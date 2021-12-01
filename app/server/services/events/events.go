package events

import (
	"context"
	"encoding/json"

	"blitzshare.api/app/dependencies"
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
	ClientId    = "bootstrap-api"
	ChannelName = "p2p-peer-registry-cmd"
)

const KubemqDefaultPort = 50000

func emitEvent(queueUrl string, event []byte, channelName string) (string, error) {
	ctx, _ := context.WithCancel(context.Background())
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(queueUrl, KubemqDefaultPort),
		kubemq.WithClientId(ClientId),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	defer client.Close()
	if err != nil {
		log.Fatalln("cant connect to queue", queueUrl, err)
	}
	sendResult, err := client.NewQueueMessage().
		SetChannel(channelName).
		SetBody(event).
		Send(ctx)
	return sendResult.MessageID, err
}

func EmitP2pPeerRegistryCmd(deps *dependencies.Dependencies, event *model.P2pPeerRegistryCmd) (string, error) {
	log.Debugln("EmitP2pPeerRegistryCmd", event)
	bEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatalln(err)
	}
	msgId, err := emitEvent(deps.Config.Settings.QueueUrl, bEvent, ChannelName)
	if err != nil {
		log.Fatalln(err)
	}
	return msgId, nil
}
