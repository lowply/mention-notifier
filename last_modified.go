package main

import (
	"io/ioutil"
	"os"
)

type LastModified struct {
	Path string
}

var lm = LastModified{
	Path: "/tmp/mention-notifier.tmp",
}

func (l *LastModified) Read() ([]byte, error) {
	_, err := os.Stat(l.Path)
	if err != nil {
		logger.Info("Creating " + l.Path + " ...")
		l.Write(nil)
	}

	lm, err := ioutil.ReadFile(l.Path)
	if err != nil {
		return nil, err
	}
	return lm, nil
}

func (l *LastModified) Write(data []byte) {
	logger.Info("Updating " + l.Path)
	ioutil.WriteFile(l.Path, data, 0644)
}
