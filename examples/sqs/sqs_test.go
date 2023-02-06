package sqs

import (
	"encoding/json"
	"fmt"
	"go-service-template/examples/sqs/helpers"
	"go-service-template/internal/configuration"
	"go-service-template/pkg/utils"
	"testing"
)

func TestSqs(t *testing.T) {
	fmt.Println(utils.Utils{}.RootPath() + "/.env")
	config, _ := configuration.Load(utils.Utils{}.RootPath() + "/.env")

	file := "/examples/sqs/samples/sqs_example.json"

	urlSqs := config.SqsHost

	messageBody, err := helpers.GetMessageRequest(file)
	if err != nil {
		fmt.Println("ERRO AO CRIAR MENSAGEM SQS NO MÉTODO GETMESSAGEREQUEST", err.Error())
	}

	createQueue := true

	existSQS := helpers.ExistSQS(config, urlSqs)

	if existSQS {
		createQueue = false
	}

	_, err = helpers.CreateSQS(createQueue, true, false, messageBody, config)
	if err != nil {
		fmt.Println("ERRO NA CRIAÇÃO DA FILA OU ENVIO DA MENSAGEM", err.Error())
	}

	messages := helpers.ReceiveMessageFromSQs(config, urlSqs)

	//messages, _ := handler.JobHandlerGetMessages(urlSqs)

	//Quando estiver testando um método real utilizar a struct da sua mensagem.
	message := make(map[string]interface{})

	if err := json.Unmarshal([]byte(messages.Body), &message); err != nil {
		fmt.Println("ERRO NA CONVERSÃO DE STRING PARA STRUCT")
	}

	helpers.DeleteMessageFromSQS(config, urlSqs)

	//Após esse ponto inserir o método que deseja testar
	fmt.Println(message)


}
