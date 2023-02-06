package helpers

import (
	"encoding/json"
	"fmt"
	"go-service-template/internal/configuration"
	awshelpersqs "go-service-template/pkg/cloud/aws"
	"go-service-template/pkg/utils"
	"io/ioutil"
	"os"
	"time"
)

// CreateSQS Metodo responsavel por criar fila SQS
func CreateSQS(createQueue bool, sendMessage bool, deleteMessage bool, messageSqs string, config *configuration.AppConfig) (string, error) {

	// Create a session instance.
	ses, err := awshelpersqs.New(awshelpersqs.Config{
		Address: config.SqsHost,
		Region:  config.RegionName,
		Profile: "localstack",
		ID:      "test",
		Secret:  "test",
	})
	if err != nil {
		return "", err
	}

	messageId, err := awshelpersqs.CreateSqsAndSendMessage(createQueue, sendMessage, deleteMessage, config, awshelpersqs.NewSQS(ses, time.Second), messageSqs)
	if err != nil {
		return "", err
	}

	return messageId, nil
}

func GetMessageRequest(file string) (string, error) {
	jsonFile, err := os.Open(utils.Utils{}.RootPath() + file)
	if err != nil {
		fmt.Println("[TEST] ERRO AO ABRIR ARQUIVO JSON")
		fmt.Println(err)
		return "", err
	}

	jsonFileValue, erro := ioutil.ReadAll(jsonFile)
	if erro != nil {
		fmt.Println("[TEST] ERRO AO LER ARQUIVO JSON")
		return "", erro
	}

	request := make(map[string]interface{})

	//Conversão da variável byte em um objeto do tipo struct sqs message
	if err := json.Unmarshal(jsonFileValue, &request); err != nil {
		fmt.Println("[TEST] ERRO UNMARCHAL MESSAGE PROTOCOL 1")
		fmt.Println(err.Error())
	}

	response, err := json.Marshal(request)
	if err != nil {
		fmt.Println("[TEST] ERRO MARCHAL RESPONSE")
		return "", err
	}

	return string(response), nil
}

func ExistSQS(config *configuration.AppConfig, urlSqs string) bool  {
	// Create a session instance.
	ses, err := awshelpersqs.New(awshelpersqs.Config{
		Address: config.SqsHost,
		Region:  config.RegionName,
		Profile: "localstack",
		ID:      "test",
		Secret:  "test",
	})
	if err != nil {
		fmt.Println("Error in create a session instance", err)
		return false
	}

	url, err := awshelpersqs.QueueARN(awshelpersqs.NewSQS(ses, time.Second), urlSqs)
	if err != nil || url == "" {
		fmt.Println("SQS Queue not exist")
		return false
	}

	return true
}

func DeleteMessageFromSQS(config *configuration.AppConfig, urlSqs string) bool  {
	// Create a session instance.
	ses, err := awshelpersqs.New(awshelpersqs.Config{
		Address: config.SqsHost,
		Region:  config.RegionName,
		Profile: "localstack",
		ID:      "test",
		Secret:  "test",
	})
	if err != nil {
		fmt.Println("Error in create a session instance", err)
		return false
	}

	awshelpersqs.DeleteMessage(awshelpersqs.NewSQS(ses, time.Second), urlSqs)

	return true
}

func ReceiveMessageFromSQs(config *configuration.AppConfig, urlSqs string) *awshelpersqs.Message {
	// Create a session instance.
	ses, err := awshelpersqs.New(awshelpersqs.Config{
		Address: config.SqsHost,
		Region:  config.RegionName,
		Profile: "localstack",
		ID:      "test",
		Secret:  "test",
	})
	if err != nil {
		fmt.Println("Error in create a session instance", err)
		return nil
	}

	message := awshelpersqs.ReceiveMessage(awshelpersqs.NewSQS(ses, time.Second), urlSqs)

	return message
}


