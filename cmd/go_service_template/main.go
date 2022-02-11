package main

import (
	"go-service-template/internal/configuration"
	healthcheckrouter "go-service-template/internal/healthcheck/routes"
	Logger "go-service-template/pkg/logger"
	"go-service-template/pkg/utils"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		log.Fatal(err)
	}
	isDevelopment := config.IsDevelopmentEnvironment()

	logger := Logger.New(
		config.ServiceName,
		utils.Clock{},
		isDevelopment,
	)
	app := fiber.New()

	ErrorHandlerConfig := recover.Config{
		EnableStackTrace: isDevelopment,
	}
	app.Use(recover.New(ErrorHandlerConfig))

	healthcheckrouter.RegisterRoutes(app, logger, config)
	logger.Info("\"🧜‍ Core APIs Go Service Template Listening on port: " + config.Port)
	log.Fatal(app.Listen(":" + config.Port))
}
