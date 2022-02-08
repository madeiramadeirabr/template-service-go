package logger_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	Logger "go-service-template/pkg/logger"
	"go-service-template/pkg/utils"
	"testing"
)

func TestLogger(t *testing.T) {
	traceId := "c29uZWdhw6fDo28gw6kgbGVnw610aW1hIGRlZmVzYQo="
	sessionId := "aW1wb3N0byDDqSByb3Vibwo"
	serviceName := "go-service-template"
	dateFixture := utils.ClockMock{}.GetCurrentTimestamp()
	logger := Logger.Logger{
		TraceID:     traceId,
		SessionID:   sessionId,
		ServiceName: serviceName,
		Clock:       utils.ClockMock{},
	}
	message := "foo"

	t.Run("Log", func(t *testing.T) {
		t.Run("Should return a log message correctly formatted without optional fields", func(t *testing.T) {
			jsonStringMessage := fmt.Sprintf(
				`{"global_event_timestamp":"%s","level":"%s","message":"%s","service_name":"%s","session_id":"%s","trace_id":"%s"}`,
				dateFixture.String(),
				Logger.LogLevelEmergency,
				message,
				serviceName,
				sessionId,
				traceId,
			)
			expectedMessage, err := utils.IndentJsonString(jsonStringMessage)
			assert.Nil(t, err, fmt.Sprintf("Expected no error, got: '%s'", err))
			loggedMessage, err := logger.Log(message, Logger.LogLevelEmergency, Logger.LogMessageOptions{})
			assert.Nil(t, err, fmt.Sprintf("Expected no error, got: '%s'", err))
			assert.Equal(t, expectedMessage, loggedMessage)
		})
		t.Run("Should return a log message correctly formatted with the optional fields set", func(t *testing.T) {
			globalEventName := "GO_SERVICE_TEMPLATE_EXAMPLE_EVENT_TOPIC"
			context := map[string]string{
				"bar": "baz",
			}
			jsonStringMessage := fmt.Sprintf(
				`{"global_event_timestamp":"%s","global_event_name":"%s","level":"%s","context":"%s","message":"%s","service_name":"%s","session_id":"%s","trace_id":"%s"}`,
				dateFixture.String(),
				globalEventName,
				Logger.LogLevelEmergency,
				fmt.Sprintf("%s", context),
				message,
				serviceName,
				sessionId,
				traceId,
			)
			expectedMessage, err := utils.IndentJsonString(jsonStringMessage)
			assert.Nil(t, err, fmt.Sprintf("Expected no error, got: '%s'", err))
			loggedMessage, err := logger.Log(
				message,
				Logger.LogLevelEmergency,
				Logger.LogMessageOptions{GlobalEventName: globalEventName, Context: context},
			)
			assert.Nil(t, err, fmt.Sprintf("Expected no error, got: '%s'", err))
			assert.Equal(t, expectedMessage, loggedMessage)
		})
	})
}
