package main

import (
	"strings"
	"time"
)

type comment struct {
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		Login     string `json:"login"`
		AvatarURL string `json:"avatar_url"`
		HTMLURL   string `json:"html_url"`
	} `json:"user"`
	Body    string `json:"body"`
	HTMLURL string `json:"html_url"`
}

func (c *comment) mentioned(login string) bool {
	if strings.Contains(c.Body, "@"+login) {
		return true
	}
	log.Debugln("There is a notification, but the latest comment didn't mention you.")
	return false
}
