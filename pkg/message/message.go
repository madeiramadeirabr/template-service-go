package message

import (
	"context"
	"go-service-template/internal/configuration"
	"go-service-template/pkg/cloud"
	"log"
)

func Message(createQueue bool, config *configuration.AppConfig, client cloud.MessageClient, messageBody string) {
	ctx := context.Background()

	queueURL := config.SqsHost
	if createQueue {
		queueURL = createQueueSQS(ctx, client)
	}
	send(ctx, client, queueURL, messageBody)
	//rcvHnd := receive(ctx, client, queURL)
	//deleteMessage(ctx, client, queURL, rcvHnd)
}

func createQueueSQS(ctx context.Context, client cloud.MessageClient) string {
	url, err := client.CreateQueue(ctx, "test-queue", false)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("create queue:", url)

	return url
}

func send(ctx context.Context, client cloud.MessageClient, queueURL string, messageBody string) {
	message := "Mensagem de Teste"
	if messageBody != "" {
		message = messageBody
	}


	id, err := client.Send(ctx, &cloud.SendRequest{
		QueueURL: queueURL,
		Body:     message,
		Attributes: []cloud.Attribute{
			{
				Key:   "Title",
				Value: "SQS send message",
				Type:  "String",
			},
			{
				Key:   "Year",
				Value: "2021",
				Type:  "Number",
			},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("send: message ID:", id)
}