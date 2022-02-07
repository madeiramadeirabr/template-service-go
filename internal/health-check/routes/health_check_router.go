package healthCheckRouter

import (
	healthCheckHandler "go-service-template/internal/health-check/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(healthCheckHandler.GetStatus())
	})
}
