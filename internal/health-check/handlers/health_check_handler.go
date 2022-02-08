package healthcheck

import "go-service-template/pkg/logger"

func GetStatus(logger *logger.Logger) map[string]interface{} {
	logger.Info("Executing Health Check...")
	return map[string]interface{}{
		"status": "OK",
	}
}
