package logger

import (
	"encoding/json"
	"fmt"
	"go-service-template/pkg/utils"
	"time"
)

type Logger struct {
	SessionID                string
	ServiceName              string
	TraceID                  string
	Clock                    utils.ClockInterface
	IsDevelopmentEnvironment bool
}

type LogLevelEnum string

const (
	LogLevelEmergency LogLevelEnum = "EMERGENCY"
	LogLevelError     LogLevelEnum = "ERROR"
	LogLevelWarn      LogLevelEnum = "WARN"
	LogLevelInfo      LogLevelEnum = "INFO"
	LogLevelDebug     LogLevelEnum = "DEBUG"
	LogLevelTrace     LogLevelEnum = "TRACE"
)

type LogMessage struct {
	GlobalEventTimestamp string       `json:"global_event_timestamp"`
	GlobalEventName      string       `json:"global_event_name,omitempty"`
	Level                LogLevelEnum `json:"level"`
	Context              string       `json:"context,omitempty"`
	Message              string       `json:"message"`
	ServiceName          string       `json:"service_name"`
	SessionID            string       `json:"session_id,omitempty"`
	TraceID              string       `json:"trace_id,omitempty"`
}

type LogMessageOptions struct {
	GlobalEventName string      `json:"global_event_name,omitempty"`
	Context         interface{} `json:"context,omitempty"`
}

func (logger Logger) formatMessage(
	message string,
	level LogLevelEnum,
	timestamp time.Time,
	logMessageOptions LogMessageOptions,
) (string, error) {
	logMessage := LogMessage{
		GlobalEventTimestamp: timestamp.String(),
		Level:                level,
		Message:              message,
		ServiceName:          logger.ServiceName,
		SessionID:            logger.SessionID,
		TraceID:              logger.TraceID,
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

func (logger Logger) Log(message string, logLevel LogLevelEnum, logMessageOptions LogMessageOptions) {
	formattedMessage, err := logger.formatMessage(
		message,
		logLevel,
		logger.Clock.GetCurrentTimestamp(),
		logMessageOptions,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(formattedMessage)
}

func (logger Logger) Emergency(message string) {
	logger.Log(message, LogLevelEmergency, LogMessageOptions{})
}

func (logger Logger) EmergencyWithOptions(message string, logMessageOptions LogMessageOptions) {
	logger.Log(message, LogLevelEmergency, logMessageOptions)
}

func (logger Logger) Error(message string) {
	logger.Log(message, LogLevelError, LogMessageOptions{})
}

func (logger Logger) ErrorWithOptions(message string, logMessageOptions LogMessageOptions) {
	logger.Log(message, LogLevelError, logMessageOptions)
}

func (logger Logger) Warn(message string) {
	logger.Log(message, LogLevelWarn, LogMessageOptions{})
}

func (logger Logger) WarnWithOptions(message string, logMessageOptions LogMessageOptions) {
	logger.Log(message, LogLevelWarn, logMessageOptions)
}

func (logger Logger) Info(message string) {
	logger.Log(message, LogLevelInfo, LogMessageOptions{})
}

func (logger Logger) InfoWithOptions(message string, logMessageOptions LogMessageOptions) {
	logger.Log(message, LogLevelInfo, logMessageOptions)
}

func (logger Logger) Debug(message string) {
	if !logger.IsDevelopmentEnvironment {
		return
	}
	logger.Log(message, LogLevelDebug, LogMessageOptions{})
}

func (logger Logger) DebugWithOptions(message string, logMessageOptions LogMessageOptions) {
	if !logger.IsDevelopmentEnvironment {
		return
	}
	logger.Log(message, LogLevelDebug, logMessageOptions)
}

func (logger Logger) Trace(message string) {
	if !logger.IsDevelopmentEnvironment {
		return
	}
	logger.Log(message, LogLevelTrace, LogMessageOptions{})
}

func (logger Logger) TraceWithOptions(message string, logMessageOptions LogMessageOptions) {
	if !logger.IsDevelopmentEnvironment {
		return
	}
	logger.Log(message, LogLevelTrace, logMessageOptions)
}
