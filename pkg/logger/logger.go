package Logger

import (
	"encoding/json"
	"fmt"
	"go-service-template/pkg/utils"
	"time"
)

type Logger struct {
	SessionId   string
	ServiceName string
	TraceId     string
	Clock       utils.ClockInterface
}

type LogLevel string

const (
	LogLevelEmergency LogLevel = "EMERGENCY"
	LogLevelError     LogLevel = "ERROR"
	LogLevelWarn      LogLevel = "WARN"
	LogLevelInfo      LogLevel = "INFO"
	LogLevelDebug     LogLevel = "DEBUG"
	LogLevelTrace     LogLevel = "TRACE"
)

type LogMessage struct {
	GlobalEventTimestamp string   `json:"global_event_timestamp"`
	GlobalEventName      string   `json:"global_event_name,omitempty"`
	Level                LogLevel `json:"level"`
	Context              string   `json:"context"`
	Message              string   `json:"message"`
	ServiceName          string   `json:"service_name"`
	SessionId            string   `json:"session_id"`
	TraceId              string   `json:"trace_id"`
}

func (logger Logger) formatMessage(
	message string,
	context string,
	level LogLevel,
	timestamp time.Time,
	globalEventName string,
) (string, error) {
	logMessage := LogMessage{
		GlobalEventTimestamp: timestamp.String(),
		Level:                level,
		Context:              context,
		Message:              message,
		ServiceName:          logger.ServiceName,
		SessionId:            logger.SessionId,
		TraceId:              logger.TraceId,
	}
	if globalEventName != "" {
		logMessage.GlobalEventName = globalEventName
	}
	formattedLogMessage, err := json.MarshalIndent(logMessage, "", "    ")
	if err != nil {
		return "", err
	}
	return string(formattedLogMessage), nil
}

func (logger Logger) Emergency(message string, context interface{}) (string, error) {
	formattedMessage, err := logger.formatMessage(
		message,
		fmt.Sprintf("%s", context),
		LogLevelEmergency,
		logger.Clock.GetCurrentTimestamp(),
		"",
	)
	if err != nil {
		return "", err
	}
	fmt.Print(formattedMessage)
	return formattedMessage, nil
}

// EmergencyEvent TODO: eventName should be a fixed enum with the available event topics
func (logger Logger) EmergencyEvent(message string, context interface{}, eventName string) (string, error) {
	formattedMessage, err := logger.formatMessage(
		message,
		fmt.Sprintf("%s", context),
		LogLevelEmergency,
		logger.Clock.GetCurrentTimestamp(),
		eventName,
	)
	if err != nil {
		return "", err
	}
	fmt.Print(formattedMessage)
	return formattedMessage, nil
}

func (logger Logger) Error(message string) {
	// TODO
}

func (logger Logger) Warn(message string) {
	// TODO
}

func (logger Logger) Info(message string) {
	// TODO
}

func (logger Logger) Debug(message string) {
	// TODO: feature toggle
}

func (logger Logger) Trace(message string) {
	// TODO: feature toggle
}
