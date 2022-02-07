package healthCheck

import (
	healthCheckHandler "go-service-template/internal/health-check/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/health-check", healthCheckHandler.GetStatus)
}
