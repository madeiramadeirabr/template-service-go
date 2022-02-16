package healthcheck_test

import (
	"go-service-template/internal/configuration"
	"go-service-template/pkg/logger"
	"testing"

	"github.com/stretchr/testify/mock"

	healthCheckHandler "go-service-template/internal/healthcheck/handlers"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	loggerMock := new(logger.Mock)
	config, _ := configuration.Load()
	t.Run("GetStatus", func(t *testing.T) {
		t.Run("It should match status message", func(t *testing.T) {
			loggerMock.On("Info", mock.Anything, mock.Anything).Return()
			status := healthCheckHandler.GetStatus(loggerMock, config)
			expected := map[string]interface{}{
				"serviceName": config.ServiceName,
				"status":      "OK",
				"environment": config.ApplicationEnv,
			}
			assert.Equal(t, status, expected)
		})
	})
}
