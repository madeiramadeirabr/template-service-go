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
	Context              string   `json:"context,omitempty"`
	Message              string   `json:"message"`
	ServiceName          string   `json:"service_name"`
	SessionId            string   `json:"session_id"`
	TraceId              string   `json:"trace_id"`
}

type LogMessageOptions struct {
	GlobalEventName string      `json:"global_event_name,omitempty"`
	Context         interface{} `json:"context,omitempty"`
}

func (logger Logger) formatMessage(
	message string,
	level LogLevel,
	timestamp time.Time,
	logMessageOptions LogMessageOptions,
) (string, error) {
	logMessage := LogMessage{
		GlobalEventTimestamp: timestamp.String(),
		Level:                level,
		Message:              message,
		ServiceName:          logger.ServiceName,
		SessionId:            logger.SessionId,
		TraceId:              logger.TraceId,
	}
	if logMessageOptions.GlobalEventName != "" {
		logMessage.GlobalEventName = logMessageOptions.GlobalEventName
	}
	if logMessageOptions.Context != nil {
		if context := fmt.Sprintf("%s", logMessageOptions.Context); context != "" {
			logMessage.Context = context
		}
	}
	formattedLogMessage, err := json.MarshalIndent(logMessage, "", "    ")
	if err != nil {
		return "", err
	}
	return string(formattedLogMessage), nil
}

func (logger Logger) Emergency(message string) (string, error) {
	return logger.EmergencyWithOptions(message, LogMessageOptions{})
}

func (logger Logger) EmergencyWithOptions(message string, logMessageOptions LogMessageOptions) (string, error) {
	formattedMessage, err := logger.formatMessage(
		message,
		LogLevelEmergency,
		logger.Clock.GetCurrentTimestamp(),
		logMessageOptions,
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
