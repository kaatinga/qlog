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

func NewStdLogger() Logger {
	return &defaultLogger{}
}

func (l *defaultLogger) log(lvl level, msg string, a ...any) {
	var output *os.File
	if lvl == LevelError {
		output = os.Stderr
	} else {
		output = os.Stdout
	}
	formatted := "[" + lvl.String() + "] " + fmt.Sprintf(msg, a...)
	output.Write([]byte(formatted))
}

func (l *defaultLogger) Debugf(msg string, a ...any) { l.log(LevelDebug, msg, a...) }
func (l *defaultLogger) Errorf(msg string, a ...any) { l.log(LevelError, msg, a...) }
