package healthcheck

import (
	"go-service-template/internal/configuration"
	healthCheckHandler "go-service-template/internal/healthcheck/handlers"
	"go-service-template/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, logger *logger.Logger, config *configuration.AppConfig) {
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(healthCheckHandler.GetStatus(logger, config))
	})

	app.Get("/alive", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(healthCheckHandler.Alive())
	})
}
