package main

import (
	"go-service-template/internal/configuration"
	healthcheckrouter "go-service-template/internal/health-check/routes"
	Logger "go-service-template/pkg/logger"
	"go-service-template/pkg/utils"

	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configuration.Load()
	utils.Fck(err)
	logger := Logger.New(
		config.ServiceName,
		utils.Clock{},
		config.IsDevelopmentEnvironment(),
	)
	app := fiber.New()
	healthcheckrouter.RegisterRoutes(app, logger, config)
	logger.Info("\"🧜‍ Core APIs Go Service Template Listening on port: " + config.Port)
	log.Fatal(app.Listen(":" + config.Port))
}
