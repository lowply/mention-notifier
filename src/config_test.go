package main

import (
	"os"
	"testing"
)

func TestDir(t *testing.T) {
	result := config.Dir()
	home := os.Getenv("HOME")
	if result != home+"/.config" {
		t.Fatalf("Failed to test: %#v", result)
	}
}

func TestLogpath(t *testing.T) {
	result := config.Logpath()
	home := os.Getenv("HOME")
	if result != home+"/.log/mention-notifier.log" {
		t.Fatalf("Failed to test: %#v", result)
	}
}
