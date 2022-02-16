package logger

import (
	"go-service-template/pkg/clock"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
	SessionID                string
	ServiceName              string
	TraceID                  string
	Clock                    clock.Interface
	IsDevelopmentEnvironment bool
}

func (loggerMock *Mock) FormatMessage(
	message string,
	level LogLevelEnum,
	logMessageOptions LogMessageOptions,
) (string, error) {
	args := loggerMock.Called(message, level, logMessageOptions)
	return args.String(0), args.Error(1)
}

func (loggerMock *Mock) Log(message string, logLevel LogLevelEnum, logMessageOptions LogMessageOptions) {
	loggerMock.Called(message, logLevel, logMessageOptions)
}

func (loggerMock *Mock) Emergency(message string) {
	loggerMock.Called(message)
}

func (loggerMock *Mock) EmergencyWithOptions(message string, logMessageOptions LogMessageOptions) {
	loggerMock.Called(message, logMessageOptions)
}

func (loggerMock *Mock) Error(message string) {
	loggerMock.Called(message)
}

func (loggerMock *Mock) ErrorWithOptions(message string, logMessageOptions LogMessageOptions) {
	loggerMock.Called(message, logMessageOptions)
}

func (loggerMock *Mock) Warn(message string) {
	loggerMock.Called(message)
}

func (loggerMock *Mock) WarnWithOptions(message string, logMessageOptions LogMessageOptions) {
	loggerMock.Called(message, logMessageOptions)
}

func (loggerMock *Mock) Info(message string) {
	loggerMock.Called(message)
}

func (loggerMock *Mock) InfoWithOptions(message string, logMessageOptions LogMessageOptions) {
	loggerMock.Called(message, logMessageOptions)
}

func (loggerMock *Mock) Debug(message string) {
	loggerMock.Called(message)
}

func (loggerMock *Mock) DebugWithOptions(message string, logMessageOptions LogMessageOptions) {
	loggerMock.Called(message, logMessageOptions)
}

func (loggerMock *Mock) Trace(message string) {
	loggerMock.Called(message)
}

func (loggerMock *Mock) TraceWithOptions(message string, logMessageOptions LogMessageOptions) {
	loggerMock.Called(message, logMessageOptions)
}
