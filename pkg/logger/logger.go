package Logger

import (
	"encoding/json"
	"go-service-template/pkg/utils"
	"time"
)

type Logger struct {
	SessionId   string
	ServiceName string
	TraceId     string
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
	GlobalEventName      string   `json:"global_event_name"`
	Level                LogLevel `json:"level"`
	Context              string   `json:"context"`
	Message              string   `json:"message"`
	ServiceName          string   `json:"service_name"`
	SessionId            string   `json:"session_id"`
	TraceId              string   `json:"trace_id"`
}

// FormatMessage TODO: see default Go logging package
func (logger Logger) FormatMessage(
	message string,
	context string,
	globalEventName string,
	level LogLevel,
	timestamp time.Time,
) string {
	logMessage := LogMessage{
		GlobalEventTimestamp: timestamp.String(),
		GlobalEventName:      globalEventName,
		Level:                level,
		Context:              context,
		Message:              message,
		ServiceName:          logger.ServiceName,
		SessionId:            logger.SessionId,
		TraceId:              logger.TraceId,
	}
	formattedLogMessage, err := json.Marshal(logMessage)
	utils.Fck(err)
	return string(formattedLogMessage)
}

func (logger Logger) Emergency(message string) {
	// TODO
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
