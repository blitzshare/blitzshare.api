package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/kubemq-io/kubemq-go"
	log "github.com/sirupsen/logrus"
)

type GenerateUploadLinkEvent struct {
	id string
}

func NewGenerateUploadLinkEvent() *GenerateUploadLinkEvent {
	return &GenerateUploadLinkEvent{id: uuid.NewString()}
}

const clientId = "fileshare-api"
const uploadMsgEventChannelName = "uploadMsgEvent"

func SubmitUploadMsgEvent(queueUrl string) *GenerateUploadLinkEvent {
	log.Info("SubmitEvent")
	event := NewGenerateUploadLinkEvent()
	ctx, _ := context.WithCancel(context.Background())

	log.Info("deps.Config.Settings.QueueUrl", queueUrl)

	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(queueUrl, 50000),
		kubemq.WithClientId(clientId),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))

	if err != nil {
		log.Fatal("something is wrong", err)
	}
	defer client.Close()
	channelName := uploadMsgEventChannelName
	meta := "some-metadata"
	bEvent, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil
	}

	err = client.E().
		SetId(event.id).
		SetChannel(channelName).
		SetMetadata(meta).
		SetBody(bEvent).
		Send(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("message sent", event.id)
	log.Info("clientId", clientId)
	log.Info("uploadMsgEventChannelName", uploadMsgEventChannelName)
	return event
}
