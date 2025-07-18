package qlog

import (
	"fmt"
	"os"
)

type defaultLogger struct{}

type level byte

const (
	LevelDebug level = iota
	LevelError
)

func (l level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func New() Logger {
	return &defaultLogger{}
}

func (l *defaultLogger) log(lvl level, msg string, args ...any) {
	var output *os.File
	if lvl == LevelError {
		output = os.Stderr
	} else {
		output = os.Stdout
	}
	formatted := "[" + lvl.String() + "] " + fmt.Sprintf(msg, args...)
	output.Write([]byte(formatted))
}

func (l *defaultLogger) Debugf(msg string, args ...any) { l.log(LevelDebug, msg, args...) }
func (l *defaultLogger) Errorf(msg string, args ...any) { l.log(LevelError, msg, args...) }
