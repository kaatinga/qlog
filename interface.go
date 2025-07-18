package qlog

type Logger interface {
	Printf(msg string, args ...any) // Printf sends a log event using debug level and no extra field.
	Errorf(msg string, args ...any) // Errorf sends a log event using error level and no extra field.
}
