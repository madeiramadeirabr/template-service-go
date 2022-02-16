package healthcheck

import (
	"go-service-template/internal/configuration"
	"go-service-template/pkg/logger"
)

// GetStatus add your health check logic here
func GetStatus(logger logger.Interface, config *configuration.AppConfig) map[string]interface{} {
	logger.Info("Executing Health Check...")
	return map[string]interface{}{
		"serviceName": config.ServiceName,
		"status":      "OK",
		"environment": config.ApplicationEnv,
	}
}

func Alive() map[string]bool {
	return map[string]bool{
		"alive": true,
	}
}
