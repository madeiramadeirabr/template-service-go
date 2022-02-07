package main

import (
	"fmt"
	configuration "go-service-template/internal/config"
	healthCheck "go-service-template/internal/health-check/routes"

	"go-service-template/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configuration.LoadConfig()
	utils.Fck(err)

	app := fiber.New()
	healthCheck.RegisterRoutes(app)

	fmt.Println("🧜‍ Core APIs Go Service Template Listening on port", config.Port)

	log.Fatal(app.Listen(":" + config.Port))
}
