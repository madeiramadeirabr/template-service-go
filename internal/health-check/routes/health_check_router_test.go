package healthcheckrouter_test

import (
	"fmt"
	healthCheckRouter "go-service-template/internal/health-check/routes"
	"go-service-template/pkg/logger"
	"go-service-template/pkg/utils"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckRouter(t *testing.T) {

	Logger := logger.New(
		"Foo",
		utils.ClockMock{},
		true,
	)

	app := fiber.New()
	healthCheckRouter.RegisterRoutes(app, Logger)

	t.Run("Health check handler", func(t *testing.T) {
		t.Run("It should return status 200", func(t *testing.T) {
			resp, err := app.Test(httptest.NewRequest("GET", "/health-check", nil))
			assert.Nil(t, err, fmt.Sprintf("Expected no error, but got '%s'", err))
			assert.Equal(t, resp.StatusCode, 200, "Expected status code 200, but got %d", resp.StatusCode)
		})
	})
}
