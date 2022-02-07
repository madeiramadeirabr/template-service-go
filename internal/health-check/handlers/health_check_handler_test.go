package healthCheck_test

import (
	"testing"

	healthCheckHandler "go-service-template/internal/health-check/handlers"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {

	t.Run("It should match status message", func(t *testing.T) {
		status := healthCheckHandler.GetStatus()
		expected := map[string]interface{}{
			"status": "OK",
		}
		assert.Equal(t, status, expected)
	})
}
