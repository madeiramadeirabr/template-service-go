package Logger_test

import (
	"bytes"
	"encoding/json"
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
		TraceId:     traceId,
		SessionId:   sessionId,
		ServiceName: serviceName,
		Clock:       utils.ClockMock{},
	}
	message := "foo"
	context := map[string]string{
		"bar": "baz",
	}

	t.Run("Emergency", func(t *testing.T) {
		t.Run("Should return an EMERGENCY log message correctly formatted", func(t *testing.T) {
			jsonStringMessage := fmt.Sprintf(
				`{"global_event_timestamp":"%s","level":"%s","context":"%s","message":"%s","service_name":"%s","session_id":"%s","trace_id":"%s"}`,
				dateFixture.String(),
				Logger.LogLevelEmergency,
				fmt.Sprintf("%s", context),
				message,
				serviceName,
				sessionId,
				traceId,
			)
			var prettyJsonStringMessage bytes.Buffer
			if err := json.Indent(&prettyJsonStringMessage, []byte(jsonStringMessage), "", "    "); err != nil {
				assert.Nil(t, err, fmt.Sprintf("Expected no error, got: '%s'", err))
			}
			expectedMessage := prettyJsonStringMessage.String()
			loggedMessage, err := logger.Emergency(message, context)
			assert.Nil(t, err, fmt.Sprintf("Expected no error, got: '%s'", err))
			assert.Equal(t, expectedMessage, loggedMessage)
		})
	})

	t.Run("EmergencyEvent", func(t *testing.T) {
		t.Run("Should return an EMERGENCY log message correctly formatted with the Global Event Name set", func(t *testing.T) {
			globalEventName := "GO_SERVICE_TEMPLATE_EXAMPLE_EVENT_TOPIC"
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
			var prettyJsonStringMessage bytes.Buffer
			if err := json.Indent(&prettyJsonStringMessage, []byte(jsonStringMessage), "", "    "); err != nil {
				assert.Nil(t, err, fmt.Sprintf("Expected no error, got: '%s'", err))
			}
			expectedMessage := prettyJsonStringMessage.String()
			loggedMessage, err := logger.EmergencyEvent(message, context, globalEventName)
			assert.Nil(t, err, fmt.Sprintf("Expected no error, got: '%s'", err))
			assert.Equal(t, expectedMessage, loggedMessage)
		})
	})

	t.Run("Error", func(t *testing.T) {
		// TODO
	})

	t.Run("Warn", func(t *testing.T) {
		// TODO
	})

	t.Run("Info", func(t *testing.T) {
		// TODO
	})

	t.Run("LogLevelDebug", func(t *testing.T) {
		// TODO
	})

	t.Run("Trace", func(t *testing.T) {
		// TODO
	})

	t.Run("Emergency", func(t *testing.T) {
		// TODO
	})
}
