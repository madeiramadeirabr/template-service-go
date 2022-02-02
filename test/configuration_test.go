package configuration_test

import (
	configuration "go-service-template/internal/config"
	"go-service-template/pkg/utils"
	"os"
	"testing"
)

// TODO: add assertion package
func TestConfiguration(t *testing.T) {
	t.Run("GetEnvString", func(t *testing.T) {
		t.Run("It should correctly get an existing string variable from the system env", func(t *testing.T) {
			err := os.Setenv("TEST_ENV_STRING", "test")
			utils.Fck(err)
			value := configuration.GetEnvString("TEST_ENV_STRING")
			if value != "test" {
				t.Errorf("Expected value to be 'test', but got '%s'", value)
			}
		})
	})

	t.Run("GetEnvString", func(t *testing.T) {
		t.Run("It should correctly get an existing boolean variable from the system env", func(t *testing.T) {
			err := os.Setenv("TEST_ENV_BOOL", "true")
			utils.Fck(err)
			value := configuration.GetEnvBool("TEST_ENV_BOOL")
			if value != true {
				t.Errorf("Expected value to be 'true', but got '%t'", value)
			}
		})
	})

	t.Run("GetEnvString", func(t *testing.T) {
		err := os.Setenv("PORT", "8080")
		utils.Fck(err)
		err = os.Setenv("DEV_ENV", "true")
		utils.Fck(err)
		config, err := configuration.LoadConfig()

		t.Run("It should return load config variables without errors", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error, but got '%s'", err)
			}
		})

		t.Run("It should return a config struct with all mandatory variables correctly filled", func(t *testing.T) {
			if config.Port != "8080" {
				t.Errorf("Expected value to be '8080', but got '%s'", config.Port)
			}
			if config.Dev != true {
				t.Errorf("Expected value to be 'true', but got '%t'", config.Dev)
			}
		})

		t.Run("It should return an error when there's a missing required variable", func(t *testing.T) {
			envErr := os.Unsetenv("PORT")
			utils.Fck(envErr)
			envErr = os.Setenv("DEV_ENV", "true")
			utils.Fck(envErr)
			_, configErr := configuration.LoadConfig()
			if configErr == nil {
				t.Errorf("Expected error, but got no error")
			}
		})
	})
}
