package logger

type Logger struct {
	TrackId string
}

type LogMessage struct {
	GlobalEventTimestamp string `json:"global_event_timestamp"`
	GlobalEventName      string `json:"global_event_name"`
	Level                string `json:"level"`
	Context              string `json:"context"`
	Message              string `json:"message"`
	ServiceName          string `json:"service_name"`
	SessionId            string `json:"session_id"`
	TraceId              string `json:"trace_id"`
}

func FormatMessage(message string, error error) string {
	// Create a struct
	// Transform into JSON String
	// return
	return "foo"
}

func Emergency(message string) {
	// TODO
}

func Error(message string) {
	// TODO
}

func Warn(message string) {
	// TODO
}

func Info(message string) {
	// TODO
}

func Debug(message string) {
	// TODO: feature toggle
}

func Trace(message string) {
	// TODO: feature toggle
}
