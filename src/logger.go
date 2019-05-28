package main

import (
	"fmt"
	"time"
)

type Logger struct {
}

var logger = Logger{}

func (_ *Logger) out(msg string) {
	t := time.Now()
	fmt.Println(t.Format(time.UnixDate) + " : " + msg)
}

func (l *Logger) Info(msg string) {
	l.out("[info] " + msg)
}

func (l *Logger) Warn(msg string) {
	l.out("[warn] " + msg)
}
