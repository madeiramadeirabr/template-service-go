package configuration

import (
	"os"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

type AppConfig struct {
	Port string `env:"PORT"`
	Dev  bool   `env:"DEV_ENV"`
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
	return validation.ValidateStruct(&c,
		validation.Field(&c.Port, validation.Required),
	)
}

func LoadConfig() (*AppConfig, error) {
	config := AppConfig{}
	config.Port = GetEnvString("PORT")
	config.Dev = GetEnvBool("DEV_ENV")
	return &config, config.Validate()
}
