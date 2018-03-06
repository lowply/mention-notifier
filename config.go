package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Config struct {
	Login          string
	GitHubToken    string
	GitHubEndpoint string
	SlackEndpoint  string
	Reason         string
	Polling        bool
}

var config = Config{
	Login:          "",
	GitHubToken:    "",
	GitHubEndpoint: "https://api.github.com/notifications",
	SlackEndpoint:  "",
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
	file, err := ioutil.ReadFile(c.Dir() + "/mention-notifier.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	if config.Login == "" || config.GitHubToken == "" || config.SlackEndpoint == "" {
		return errors.New("Invalid config")
	}

	return nil
}
