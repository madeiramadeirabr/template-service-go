package main

import (
	"fmt"
	configuration "go-service-template/internal/config"
	"go-service-template/pkg/utils"
)

func main() {
	config, err := configuration.LoadConfig()
	utils.Fck(err)
	fmt.Println("🧜‍ Core APIs Go Service Template Listening on port ", config.Port)
}
