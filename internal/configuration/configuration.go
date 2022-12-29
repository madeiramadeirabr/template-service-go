package configuration

import (
	"fmt"
	"os"
	"strconv"

	"github.com/thoas/go-funk"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/joho/godotenv"
)

type ApplicationEnvEnum string

const (
	Development ApplicationEnvEnum = "DEVELOPMENT"
	Test        ApplicationEnvEnum = "TEST"
	Production  ApplicationEnvEnum = "PRODUCTION"
	Staging     ApplicationEnvEnum = "STAGING"
)

type AppConfig struct {
	Port                string `env:"PORT"`
	ApplicationEnv      string `env:"APPLICATION_ENV"`
	IsTesting           bool   `env:"TESTING"`
	ServiceName         string `env:"SERVICE_NAME"`
	LocalStackHost      string `env:"LOCALSTACK_HOST"`
	LocalStackSqsHost   string `env:"LOCALSTACK_SQS_HOST"`
	LocalStackSqsName   string `env:"LOCALSTACK_SQS_NAME"`
	LocalStackSqsRegion string `env:"LOCALSTACK_SQS_REGION"`
}

func (appConfig AppConfig) IsDevelopmentEnvironment() bool {
	developmentEnvironments := []ApplicationEnvEnum{Test, Development, Staging}
	return funk.Contains(developmentEnvironments, ApplicationEnvEnum(appConfig.ApplicationEnv))
}

func (appConfig AppConfig) Validate() error {
	return validation.ValidateStruct(
		&appConfig,
		validation.Field(&appConfig.Port, validation.Required),
		validation.Field(&appConfig.ApplicationEnv, validation.Required),
	)
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

func Load(dotenvfiles ...string) (*AppConfig, error) {
	config := AppConfig{}
	err := godotenv.Load(dotenvfiles...)
	if err != nil {
		fmt.Println("[WARN] .env file not found. Loading from system environment")
	}
	config.Port = GetEnvString("PORT")
	config.ApplicationEnv = GetEnvString("APPLICATION_ENV")
	config.IsTesting = GetEnvBool("TESTING")
	config.ServiceName = GetEnvString("SERVICE_NAME")
	config.LocalStackHost = GetEnvString("LOCALSTACK_HOST")
	config.LocalStackSqsHost = GetEnvString("LOCALSTACK_SQS_HOST")
	config.LocalStackSqsName = GetEnvString("LOCALSTACK_SQS_NAME")
	config.LocalStackSqsRegion = GetEnvString("LOCALSTACK_SQS_REGION")
	return &config, config.Validate()
}
