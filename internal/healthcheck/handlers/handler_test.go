package healthcheck_test

import (
	"go-service-template/internal/configuration"
	"go-service-template/pkg/logger"
	"go-service-template/pkg/utils"
	"testing"

	healthCheckHandler "go-service-template/internal/healthcheck/handlers"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	Logger := logger.New(
		"Foo",
		utils.ClockMock{},
		true,
	)
	config, _ := configuration.Load()
	t.Run("GetStatus", func(t *testing.T) {
		t.Run("It should match status message", func(t *testing.T) {
			status := healthCheckHandler.GetStatus(Logger, config)
			expected := map[string]interface{}{
				"serviceName": config.ServiceName,
				"status":      "OK",
				"environment": config.ApplicationEnv,
			}
			assert.Equal(t, status, expected)
		})
	})
}
