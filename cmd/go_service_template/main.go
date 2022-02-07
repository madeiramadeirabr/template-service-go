package main

import (
	"fmt"
	"go-service-template/internal/configuration"
	healthCheck "go-service-template/internal/health-check/routes"

	"go-service-template/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configuration.Load()
	utils.Fck(err)

	app := fiber.New()
	healthCheck.RegisterRoutes(app)

	fmt.Println("🧜‍ Core APIs Go Service Template Listening on port", config.Port)

	log.Fatal(app.Listen(":" + config.Port))
}
