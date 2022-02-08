package healthcheckrouter

import (
	healthCheckHandler "go-service-template/internal/health-check/handlers"
	"go-service-template/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, logger *logger.Logger) {
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(healthCheckHandler.GetStatus(logger))
	})
}
