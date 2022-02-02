package configuration_test

import (
	configuration "go-service-template/internal/config"
	"os"
	"testing"
)

func TestGetEnvString(t *testing.T) {
	//given
	os.Setenv("TEST_ENV_STRING", "test")
	//when
	value := configuration.GetEnvString("TEST_ENV_STRING")
	//then
	if value != "test" {
		t.Errorf("Expected value to be 'test', but got '%s'", value)
	}
}

func TestGetEnvBool(t *testing.T) {
	//given
	os.Setenv("TEST_ENV_BOOL", "true")
	//when
	value := configuration.GetEnvBool("TEST_ENV_BOOL")
	//then
	if value != true {
		t.Errorf("Expected value to be 'true', but got '%t'", value)
	}
}

func TestLoadEnvWithNoError(t *testing.T) {
	//given
	os.Setenv("PORT", "8080")
	os.Setenv("DEV_ENV", "true")
	//when
	config, error := configuration.LoadConfig()
	if error != nil {
		t.Errorf("Expected no error, but got '%s'", error)
	}
	//then
	if config.Port != "8080" {
		t.Errorf("Expected value to be '8080', but got '%s'", config.Port)
	}
	if config.Dev != true {
		t.Errorf("Expected value to be 'true', but got '%t'", config.Dev)
	}
}

func TestLoadEnvWithMissingPort(t *testing.T) {
	//given
	os.Unsetenv("PORT")
	os.Setenv("DEV_ENV", "true")
	//when
	_, error := configuration.LoadConfig()
	if error == nil {
		t.Errorf("Expected error, but got no error")
	}
}
