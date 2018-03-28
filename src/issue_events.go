package main

import (
	"time"
)

type IssueEvent struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Actor struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"actor"`
	Event     string      `json:"event"`
	CommitID  interface{} `json:"commit_id"`
	CommitURL interface{} `json:"commit_url"`
	CreatedAt time.Time   `json:"created_at"`
}

type IssueEvents []IssueEvent

func (es *IssueEvents) query(url string) error {
	var r = new(Requester)
	err := r.GetAndUnmarshal(url, es)
	if err != nil {
		return err
	}
	return nil
}

func (es *IssueEvents) closedOrReopened() bool {
	for i := len(*es) - 1; i >= 0; i-- {
		e := (*es)[i].Event
		if e == "closed" || e == "reopened" {
			return true
		}
	}
	return false
}
