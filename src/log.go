package main

import (
	_log "log"
	"os"
)

type Logger struct {
	_log.Logger
	debug bool
}

func newLogger() *Logger {
	l := &Logger{}
	l.SetOutput(os.Stdout)
	l.debug = false

	if os.Getenv("DEBUG") == "true" {
		l.debug = true
	}
	return l
}

func (l *Logger) Debugln(msg string) {
	if l.debug {
		_log.Println(msg)
	}
}
