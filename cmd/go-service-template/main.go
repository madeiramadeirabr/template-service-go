package main

import (
	"fmt"
	"go-service-template/internal/configuration"
	"go-service-template/pkg/utils"
)

func main() {
	config, err := configuration.Load()
	utils.Fck(err)
	fmt.Println("🧜‍ Core APIs Go Service Template Listening on port ", config.Port)
}
