package qlog

type Logger interface {
	Debugf(msg string, a ...any)
	Errorf(msg string, a ...any)
}
