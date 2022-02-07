package Logger_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	Logger "go-service-template/pkg/logger"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	traceId := "c29uZWdhw6fDo28gw6kgbGVnw610aW1hIGRlZmVzYQo="
	sessionId := "aW1wb3N0byDDqSByb3Vibwo"
	serviceName := "go-service-template"
	dateFixture := time.Now()
	logger := Logger.Logger{
		TraceId:     traceId,
		SessionId:   sessionId,
		ServiceName: serviceName,
	}

	t.Run("FormatMessage", func(t *testing.T) {
		message := "I ain't the sharpest tool in the shed"
		context := "{ message: \"Somebody once told me the world is gonna roll me\" }"
		globalEventName := "GO_SERVICE_TEMPLATE_EXAMPLE_EVENT_TOPIC"
		expectedFormattedMessage := Logger.LogMessage{
			GlobalEventTimestamp: dateFixture.String(),
			GlobalEventName:      globalEventName,
			Level:                Logger.LogLevelDebug,
			Context:              context,
			Message:              message,
			ServiceName:          serviceName,
			SessionId:            sessionId,
			TraceId:              traceId,
		}
		formattedMessage := logger.FormatMessage(message, context, globalEventName, Logger.LogLevelDebug, dateFixture)
		expectedFormattedMessageString, err := json.Marshal(expectedFormattedMessage)
		assert.Nil(t, err, fmt.Sprintf("Expected no error, but got '%s'", err))
		assert.Equal(t, formattedMessage, string(expectedFormattedMessageString))
	})

	t.Run("Emergency", func(t *testing.T) {
		// TODO
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
