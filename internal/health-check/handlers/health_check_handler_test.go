package healthcheck_test

import (
	"go-service-template/pkg/logger"
	"go-service-template/pkg/utils"
	"testing"

	healthCheckHandler "go-service-template/internal/health-check/handlers"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	Logger := logger.New(
		"Foo",
		utils.ClockMock{},
		true,
	)
	t.Run("GetStatus", func(t *testing.T) {
		t.Run("It should match status message", func(t *testing.T) {
			status := healthCheckHandler.GetStatus(Logger)
			expected := map[string]interface{}{
				"status": "OK",
			}
			assert.Equal(t, status, expected)
		})
	})
}
