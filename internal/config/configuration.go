package configuration

import (
	"fmt"
	"os"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port           string `env:"PORT"`
	ApplicationEnv string `env:"APPLICATION_ENV"`
	IsTesting      bool   `env:"TESTING"`
}

func GetEnvString(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}

func GetEnvBool(key string) bool {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseBool(value); err == nil {
			return value
		}
	}
	return false
}

func (c AppConfig) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Port, validation.Required),
	)
}

func LoadConfig() (*AppConfig, error) {
	config := AppConfig{}
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	config.Port = GetEnvString("PORT")
	config.ApplicationEnv = GetEnvString("APPLICATION_ENV")
	config.IsTesting = GetEnvBool("TESTING")
	return &config, config.Validate()
}
