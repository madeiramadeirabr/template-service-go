package localstack

import (
	"encoding/json"
	"fmt"
	"go-service-template/internal/configuration"
	"go-service-template/pkg/cloud/aws"
	"go-service-template/pkg/message"
	"go-service-template/pkg/utils"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func LoadLocalStack(sqs, s3 bool, messageSqs string, config *configuration.AppConfig) {

	if sqs {
		// Create a session instance.
		ses, err := aws.New(aws.Config{
			Address: config.LocalStackHost,
			Region:  config.LocalStackSqsRegion,
			Profile: "localstack",
			ID:      "test",
			Secret:  "test",
		})
		if err != nil {
			log.Fatalln(err)
		}

		message.Message(aws.NewSQS(ses, time.Second), messageSqs)
	}

	if s3 {
		fmt.Println("Ainda preciso desenvolver essa parte.")
	}
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
