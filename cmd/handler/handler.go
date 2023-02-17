package handler

import (
	"errors"
	"go-service-template/internal/configuration"
	"strconv"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var sqsInstance *sqs.SQS
var sessionAws *session.Session
var once sync.Once
var config, _ = configuration.Load()

func loadSession() {
	if config.IsDevelopmentEnvironment() {
		cfg := aws.Config{
			Region:   aws.String(config.RegionName),
			Endpoint: aws.String( config.SqsHost),
		}

		once.Do(func() {
			sessionAws := session.Must(session.NewSession(&cfg))
			sqsInstance = sqs.New(sessionAws)
		})
	} else {

		once.Do(func() {
			sessionAws, _ = session.NewSessionWithOptions(session.Options{
				SharedConfigState: session.SharedConfigEnable,
			})
			sqsInstance = sqs.New(sessionAws)
		})
	}
}

func JobHandlerGetMessages(queueURL string) (*sqs.ReceiveMessageOutput, error) {
	loadSession()

	receive := sqs.ReceiveMessageInput{
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			aws.String(sqs.MessageSystemAttributeNameApproximateReceiveCount),
		},
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(7),
		VisibilityTimeout:   aws.Int64(30),
		WaitTimeSeconds:     aws.Int64(2),
	}

	result, err := sqsInstance.ReceiveMessage(&receive)

	if err != nil {
		return nil, err
	}
	return result, err
}

func JobHandlerSendMessage(queueURL string, message string, attributes map[string]*sqs.MessageAttributeValue) error {
	if strings.Index(queueURL, ".fifo") > 0 {
		return errors.New("error: This Queue is a FIFO")
	}

	loadSession()

	_, err := sqsInstance.SendMessage(&sqs.SendMessageInput{
		MessageBody:       aws.String(message),
		QueueUrl:          aws.String(queueURL),
		MessageAttributes: attributes,
	})

	return err
}

func JobHandlerDeleteMessage(queueURL string, msg string) error {
	loadSession()

	_, err := sqsInstance.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(msg),
	})

	return err
}

func GetAttributesMessage(msg sqs.Message) map[string]*sqs.MessageAttributeValue {

	var MessageAttributes map[string]*sqs.MessageAttributeValue

	if msg.MessageAttributes["Fail"] != nil {
		count := *msg.MessageAttributes["Fail"]
		countInt, _ := strconv.Atoi(*count.StringValue)
		totalCount := strconv.Itoa(countInt + 1)
		msg.MessageAttributes["Fail"].SetStringValue(totalCount)

		return msg.MessageAttributes
	}

	MessageAttributes = map[string]*sqs.MessageAttributeValue{
		"Fail": &sqs.MessageAttributeValue{
			DataType:    aws.String("Number"),
			StringValue: aws.String("1"),
		},
	}
	return MessageAttributes
}
