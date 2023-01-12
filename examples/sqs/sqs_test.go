package sqs

import (
	"encoding/json"
	"fmt"
	"go-service-template/examples/sqs/helpers"
	"go-service-template/internal/configuration"
	"go-service-template/pkg/handler"
	"go-service-template/pkg/utils"
	"testing"
)

func TestSqs(t *testing.T) {
	fmt.Println(utils.Utils{}.RootPath() + "/.env")
	config, _ := configuration.Load(utils.Utils{}.RootPath() + "/.env")

	file := "examples/sqs/samples/sqs_example.json"

	messageBody, err := helpers.GetMessageRequest(file)
	if err != nil {
		fmt.Println("ERRO AO CRIAR MENSAGEM SQS NO MÉTODO GETMESSAGEREQUEST")
	}

	helpers.CreateSQS(true, messageBody, config)

	urlSqs := config.SqsHost

	messages, _ := handler.JobHandlerGetMessages(urlSqs)

	//Quando estiver testando um método real utilizar a struct da sua mensagem.
	message := make(map[string]interface{})

	if err := json.Unmarshal([]byte(*messages.Messages[0].Body), &message); err != nil {
		fmt.Println("ERRO NA CONVERSÃO DE STRING PARA STUCT")
	}

	//Após esse ponto inserir o método que deseja testar
	fmt.Println(message)
}
