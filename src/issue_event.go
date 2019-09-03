package main

import (
	"time"
)

type issueEvent struct {
	Event     string    `json:"event"`
	CreatedAt time.Time `json:"created_at"`
}

func (i *issueEvent) closedOrReopened() bool {
	if i.Event == "closed" || i.Event == "reopened" {
		log.Debugln("This issue was just " + i.Event + ", not sending Slack notification")
		return true
	}
	return false
}
