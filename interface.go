package qlog

type Logger interface {
	Printf(msg string, args ...any)
	Errorf(msg string, args ...any)
}
