package main

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Login          string
	GitHubToken    string
	SlackEndpoint  string
	GitHubEndpoint string
	Reason         string
	Polling        bool
}

var config = Config{
	Login:          "",
	GitHubToken:    "",
	SlackEndpoint:  "",
	GitHubEndpoint: "https://api.github.com/notifications",
	Reason:         "mention",
	Polling:        false,
}

func (c *Config) Dir() string {
	return os.Getenv("HOME") + "/.config"
}

func (c *Config) Logpath() string {
	return os.Getenv("HOME") + "/.log/mention-notifier.log"
}

func (c *Config) Read() error {
	// Required
	for _, v := range []string{"LOGIN", "GITHUB_TOKEN", "SLACK_ENDPOINT"} {
		if os.Getenv(v) == "" {
			return errors.New(v + " is empty.")
		}
	}

	c.Login = os.Getenv("LOGIN")
	c.GitHubToken = os.Getenv("GITHUB_TOKEN")
	c.SlackEndpoint = os.Getenv("SLACK_ENDPOINT")

	// Options
	if os.Getenv("GITHUB_ENDPOINT") != "" {
		c.GitHubEndpoint = os.Getenv("GITHUB_ENDPOINT")
	}

	if os.Getenv("REASON") != "" {
		c.Reason = os.Getenv("REASON")
	}

	if os.Getenv("POLLING") != "" {
		b, err := strconv.ParseBool(os.Getenv("POLLING"))
		if err != nil {
			return err
		}
		c.Polling = b
	}

	return nil
}
