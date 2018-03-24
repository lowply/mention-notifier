package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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

func (es *IssueEvents) Get(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "token "+config.GitHubToken)
	logger.Info("GET " + url)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	logger.Info("DONE " + resp.Status)

	if resp.StatusCode != 200 {
		return errors.New("Unable to access to the endpoint: " + url)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &es)
	if err != nil {
		return err
	}

	return nil
}

func (es IssueEvents) ClosedOrReopened() bool {
	for i := len(es) - 1; i >= 0; i-- {
		e := es[i].Event
		if e == "reopened" || e == "closed" {
			return true
		}
	}
	return false
}
