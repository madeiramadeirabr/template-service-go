package main

import (
	"go-service-template/internal/configuration"
	healthCheckRouter "go-service-template/internal/health-check/routes"
	"go-service-template/pkg/logger"
	"go-service-template/pkg/utils"

	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configuration.Load()
	utils.Fck(err)
	Logger := logger.Logger{
		ServiceName:              config.ServiceName,
		Clock:                    utils.Clock{},
		IsDevelopmentEnvironment: config.IsDevelopmentEnvironment(),
	}
	app := fiber.New()
	healthCheckRouter.RegisterRoutes(app, &Logger)
	Logger.Info("\"🧜‍ Core APIs Go Service Template Listening on port: " + config.Port)
	log.Fatal(app.Listen(":" + config.Port))
}
