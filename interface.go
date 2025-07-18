package qlog

type Logger interface {
	Debugf(msg string, args ...any)
	Errorf(msg string, args ...any)
}
