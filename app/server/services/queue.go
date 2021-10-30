package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	kubemq "github.com/kubemq-io/kubemq-go"
	log "github.com/sirupsen/logrus"
)

type GenerateUploadLinkEvent struct {
	Id string `json:"id"`
}

func NewGenerateUploadLinkEvent() *GenerateUploadLinkEvent {
	return &GenerateUploadLinkEvent{Id: uuid.NewString()}
}

const clientId = "fileshare-api"
const uploadMsgEventChannelName = "uploadMsgEventChaannel"

func SubmitUploadMsgEvent(queueUrl string) *GenerateUploadLinkEvent {
	log.Info("SubmitEvent")
	event := NewGenerateUploadLinkEvent()
	ctx, _ := context.WithCancel(context.Background())
	log.Info("deps.Config.Settings.QueueUrl", queueUrl)
	bEvent, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil
	}
	log.Info("bEvent", bEvent)
	log.Info("str:bEvent", string(bEvent))
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

	err = client.E().
		SetId(event.Id).
		SetChannel(channelName).
		SetMetadata(meta).
		SetBody(bEvent).
		Send(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("message sent", event.Id)
	log.Info("clientId", clientId)
	log.Info("uploadMsgEventChannelName", uploadMsgEventChannelName)
	return event
}
