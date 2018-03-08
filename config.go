package main

import (
	"errors"
	"os"
)

type Config struct {
	Login          string
	GitHubToken    string
	SlackEndpoint  string
	GitHubEndpoint string
	Reason         string
	Polling        string
}

var config = Config{
	Login:          "",
	GitHubToken:    "",
	SlackEndpoint:  "",
	GitHubEndpoint: "https://api.github.com/notifications",
	Reason:         "mention",
	Polling:        "false",
}

func (c *Config) Dir() string {
	return os.Getenv("HOME") + "/.config"
}

func (c *Config) Logpath() string {
	return os.Getenv("HOME") + "/.log/mention-notifier.log"
}

func (c *Config) Read() error {
	// Required
	if os.Getenv("LOGIN") == "" {
		return errors.New("LOGIN " + "is empty.")
	}
	c.Login = os.Getenv("LOGIN")

	if os.Getenv("GITHUB_TOKEN") == "" {
		return errors.New("GITHUB_TOKEN" + " " + "is empty.")
	}
	c.GitHubToken = os.Getenv("GITHUB_TOKEN")

	if os.Getenv("SLACK_ENDPOINT") == "" {
		return errors.New("SLACK_ENDPOINT" + " " + "is empty.")
	}
	c.SlackEndpoint = os.Getenv("SLACK_ENDPOINT")

	// Options
	if os.Getenv("GITHUB_ENDPOINT") != "" {
		c.GitHubEndpoint = os.Getenv("GITHUB_ENDPOINT")
	}

	if os.Getenv("REASON") != "" {
		c.Reason = os.Getenv("REASON")
	}

	if os.Getenv("POLLING") != "" {
		c.Polling = os.Getenv("POLLING")
	}

	return nil
}
