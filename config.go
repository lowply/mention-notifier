package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
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
	return os.Getenv("HOME") + "/" + ".config"
}

func (c *Config) CheckDir() error {
	_, err := os.Stat(c.Dir())
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			err := os.Mkdir(config.Dir(), 0755)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Config) Read() error {
	err := c.CheckDir()
	if err != nil {
		return err
	}

	path := c.Dir() + "/mention-notifier.json"
	file, err := ioutil.ReadFile(path)
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
