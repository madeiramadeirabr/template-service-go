package aws

import (
	"context"
	"go-service-template/internal/configuration"
	"log"
)

func CreateSqsAndSendMessage(createQueue bool, sendMessage bool, config *configuration.AppConfig, client MessageClient, messageBody string) (string, error) {
	ctx := context.Background()

	queueURL := config.SqsHost
	if createQueue {
		//Cria a fila SQS
		queueURL = createQueueSQS(ctx, client)
	}

	if sendMessage {
		//Envia a mensagem para a fila criada.
		id, err := sendMessageToSqs(ctx, client, queueURL, messageBody)
		if err != nil {
			return "", err
		}

		return id, nil
	}

	return "", nil
}

func ReceiveMessage(client MessageClient, queueURL string) *Message {
	ctx := context.Background()

	message := receive(ctx, client, queueURL)

	return message
}

func DeleteMessage(client MessageClient, queueURL string) {
	ctx := context.Background()

	rcvHnd := receive(ctx, client, queueURL)

	//Remove a mensagem da fila
	deleteMessageFromSQS(ctx, client, queueURL, rcvHnd.ReceiptHandle)
}

func createQueueSQS(ctx context.Context, client MessageClient) string {
	url, err := client.CreateQueue(ctx, "test-queue", false)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("create queue:", url)

	return url
}

func sendMessageToSqs(ctx context.Context, client MessageClient, queueURL string, messageBody string) (string, error) {
	message := "Mensagem de Teste"
	if messageBody != "" {
		message = messageBody
	}

	id, err := client.SendMessage(ctx, &SendRequest{
		QueueURL: queueURL,
		Body:     message,
		Attributes: []Attribute{
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
		return "", err
	}
	return id, nil
}

func receive(ctx context.Context, client MessageClient, queueURL string) *Message {
	res, err := client.ReceiveMessage(ctx, queueURL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("receive:", res)

	return res
}

func deleteMessageFromSQS(ctx context.Context, client MessageClient, queueURL, rcvHnd string) {
	if err := client.DeleteMessage(ctx, queueURL, rcvHnd); err != nil {
		log.Fatalln(err)
	}
	log.Println("delete message: ok")
}

func QueueARN(client MessageClient, url string) (string, error) {
	ctx := context.Background()
	arn, err := client.QueueARN(ctx, url)
	if err != nil {
		return "", err
	}

	return arn, nil
}