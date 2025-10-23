package logger

import (
	"log"
	"os"
)

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type stdLogger struct {
	base *log.Logger
}

func New() Logger {
	return &stdLogger{base: log.New(os.Stdout, "", log.LstdFlags)}
}

func (l *stdLogger) Infof(format string, args ...interface{}) {
	l.base.Printf("INFO: "+format, args...)
}

func (l *stdLogger) Errorf(format string, args ...interface{}) {
	l.base.Printf("ERROR: "+format, args...)
}
