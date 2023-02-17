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
	SqsEndpoint		    string `env:"SQS_ENDPOINT"`
	SqsHost   			string `env:"SQS_HOST"`
	SqsName   			string `env:"SQS_NAME"`
	RegionName 			string `env:"REGION_NAME"`
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
	config.SqsEndpoint = GetEnvString("SQS_ENDPOINT")
	config.SqsHost = GetEnvString("SQS_HOST")
	config.SqsName = GetEnvString("SQS_NAME")
	config.RegionName = GetEnvString("REGION_NAME")
	return &config, config.Validate()
}
