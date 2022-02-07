package configuration_test

import (
	"fmt"
	"go-service-template/internal/configuration"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	t.Run("GetEnvString", func(t *testing.T) {
		t.Run("It should correctly get an existing string variable from the system env", func(t *testing.T) {
			_ = os.Setenv("TEST_ENV_STRING", "test")
			value := configuration.GetEnvString("TEST_ENV_STRING")
			assert.Equal(t, value, "test", fmt.Sprintf("Expected value to be 'test', but got '%s'", value))
		})
		defer func() {
			_ = os.Unsetenv("TEST_ENV_STRING")
		}()
	})

	t.Run("GetEnvBool", func(t *testing.T) {
		t.Run("It should correctly get an existing boolean variable from the system env", func(t *testing.T) {
			_ = os.Setenv("TEST_ENV_BOOL", "true")
			value := configuration.GetEnvBool("TEST_ENV_BOOL")
			assert.Equal(t, value, true, fmt.Sprintf("Expected value to be 'true', but got '%t'", value))
		})
		defer func() {
			_ = os.Unsetenv("TEST_ENV_BOOL")
		}()
	})

	t.Run("Load", func(t *testing.T) {
		_ = os.Setenv("PORT", "8080")
		_ = os.Setenv("APPLICATION_ENV", "TEST")
		_ = os.Setenv("TESTING", "true")
		_ = os.Setenv("APPLICATION_ENV", "TEST")

		config, err := configuration.Load()

		t.Run("It should return load configuration variables without errors", func(t *testing.T) {
			assert.Nil(t, err, fmt.Sprintf("Expected no error, but got '%s'", err))
		})

		t.Run("It should return a configuration struct with all mandatory variables correctly filled", func(t *testing.T) {
			assert.Equal(t, config.Port, "8080", fmt.Sprintf("Expected value to be '8080', but got '%s'", config.Port))
			assert.Equal(t, config.IsTesting, true, fmt.Sprintf("Expected value to be 'true', but got '%t'", config.IsTesting))
		})

		t.Run("It should return an error when there's a missing required variable", func(t *testing.T) {
			_ = os.Unsetenv("PORT")
			_, configErr := configuration.Load()
			assert.NotNil(t, configErr, "Expected error, but got nil")
		})

		defer func() {
			_ = os.Unsetenv("PORT")
			_ = os.Unsetenv("TESTING")
			_ = os.Unsetenv("APPLICATION_ENV")
		}()
	})
}
