package events

import (
	"context"
	"encoding/json"

	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	kubemq "github.com/kubemq-io/kubemq-go"
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

func emitEvent(queueUrl string, event []byte, channelName string) (string, error) {
	ctx, _ := context.WithCancel(context.Background())
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(queueUrl, 50000),
		kubemq.WithClientId(ClientId),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))

	if err != nil {
		log.Errorln("cant connect to queue", queueUrl, err)
		return "", err
	}
	defer client.Close()

	sendResult, err := client.NewQueueMessage().
		SetChannel(channelName).
		SetBody(event).
		Send(ctx)

	if err != nil {
		return "", err
	}
	log.Debugln("clientId", ClientId)
	log.Debugln("uploadMsgEventChannelName", ChannelName)
	return sendResult.MessageID, nil
}

func EmitP2pPeerRegistyCmd(deps *dependencies.Dependencies, event *model.P2pPeerRegistryCmd) (string, error) {
	log.Debugln("EmitP2pPeerRegistryCmd", event)
	bEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	msgId, err := emitEvent(deps.Config.Settings.QueueUrl, bEvent, ChannelName)
	if err != nil {
		return "", err
	}
	log.Debugln("msgId", msgId)
	return msgId, nil
}
