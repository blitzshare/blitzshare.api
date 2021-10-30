package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	kubemq "github.com/kubemq-io/kubemq-go"
	log "github.com/sirupsen/logrus"
)

type GenerateUploadLinkEvent struct {
	ShareId string `json:"shareId"`
}

func NewGenerateUploadLinkEvent() *GenerateUploadLinkEvent {
	return &GenerateUploadLinkEvent{ShareId: uuid.NewString()}
}

const clientId = "fileshare-api"
const uploadMsgEventChannelName = "uploadMsgEventChaannel"

func SubmitUploadMsgEvent(queueUrl string) string {
	log.Info("SubmitEvent:Queue")
	event := NewGenerateUploadLinkEvent()
	ctx, _ := context.WithCancel(context.Background())
	log.Info("deps.Config.Settings.QueueUrl", queueUrl)
	bEvent, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(queueUrl, 50000),
		kubemq.WithClientId(clientId),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))

	if err != nil {
		log.Fatal("something is wrong", err)
	}
	defer client.Close()

	sendResult, err := client.NewQueueMessage().
		SetChannel(uploadMsgEventChannelName).
		SetBody(bEvent).
		Send(ctx)

	if err != nil {
		log.Fatal(err)
	}
	log.Infoln("clientId", clientId)
	log.Infoln("uploadMsgEventChannelName", uploadMsgEventChannelName)
	return sendResult.MessageID
}

func SubscribeToQueue(queueUrl string) {
	fmt.Println("Infinite Loop 2")
	log.Info("SubmitEvent:SubscribeToQueue")
	for true {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		client, err := kubemq.NewClient(ctx,
			kubemq.WithAddress(queueUrl, 50000),
			kubemq.WithClientId(clientId),
			kubemq.WithTransportType(kubemq.TransportTypeGRPC))
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		receiveResult, err := client.NewReceiveQueueMessagesRequest().
			SetChannel(uploadMsgEventChannelName).
			SetMaxNumberOfMessages(1).
			SetWaitTimeSeconds(5).
			Send(ctx)
		if err != nil {
			log.Fatal(err)
		}
		if receiveResult.Messages != nil {
			log.Printf("Received %d Messages:\n", receiveResult.MessagesReceived)
			for _, msg := range receiveResult.Messages {
				log.Printf("MessageID: %s, Body: %s", msg.MessageID, string(msg.Body))
			}
		}
		time.Sleep(time.Second * 10)
	}
}
