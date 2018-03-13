package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type Logger struct {
}

var logger = Logger{}

func (_ *Logger) outfile(msg string) {
	_ = os.Mkdir(path.Dir(config.Logpath()), 0755)

	f, err := os.OpenFile(config.Logpath(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	t := time.Now()
	f.Write([]byte(t.Format(time.UnixDate) + " : " + msg + "\n"))
}

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
