package configuration_test

import (
	"fmt"
	configuration "go-service-template/internal/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	t.Run("GetEnvString", func(t *testing.T) {
		t.Run("It should correctly get an existing string variable from the system env", func(t *testing.T) {
			os.Setenv("TEST_ENV_STRING", "test")
			value := configuration.GetEnvString("TEST_ENV_STRING")
			assert.Equal(t, value, "test", fmt.Sprintf("Expected value to be 'test', but got '%s'", value))
		})

		defer os.Unsetenv("TEST_ENV_STRING")
	})

	t.Run("GetEnvString", func(t *testing.T) {
		t.Run("It should correctly get an existing boolean variable from the system env", func(t *testing.T) {
			os.Setenv("TEST_ENV_BOOL", "true")
			value := configuration.GetEnvBool("TEST_ENV_BOOL")
			assert.Equal(t, value, true, fmt.Sprintf("Expected value to be 'true', but got '%t'", value))
		})

		defer os.Unsetenv("TEST_ENV_BOOL")
	})

	t.Run("GetEnvString", func(t *testing.T) {
		os.Setenv("PORT", "8080")
		os.Setenv("DEV_ENV", "true")

		config, err := configuration.LoadConfig()

		t.Run("It should return load config variables without errors", func(t *testing.T) {
			assert.Nil(t, err, fmt.Sprintf("Expected no error, but got '%s'", err))
		})

		t.Run("It should return a config struct with all mandatory variables correctly filled", func(t *testing.T) {
			assert.Equal(t, config.Port, "8080", fmt.Sprintf("Expected value to be '8080', but got '%s'", config.Port))
			assert.Equal(t, config.Dev, true, fmt.Sprintf("Expected value to be 'true', but got '%t'", config.Dev))
		})

		t.Run("It should return an error when there's a missing required variable", func(t *testing.T) {
			os.Unsetenv("PORT")
			_, configErr := configuration.LoadConfig()
			assert.NotNil(t, configErr, "Expected error, but got nil")
		})

		defer os.Unsetenv("PORT")
		defer os.Unsetenv("DEV_ENV")
	})
}
