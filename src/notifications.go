package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Notification struct {
	ID         string `json:"id"`
	Repository struct {
		ID    int `json:"id"`
		Owner struct {
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
		} `json:"owner"`
		Name        string `json:"name"`
		FullName    string `json:"full_name"`
		Description string `json:"description"`
		Private     bool   `json:"private"`
		Fork        bool   `json:"fork"`
		URL         string `json:"url"`
		HTMLURL     string `json:"html_url"`
	} `json:"repository"`
	Subject struct {
		Title            string `json:"title"`
		URL              string `json:"url"`
		LatestCommentURL string `json:"latest_comment_url"`
		Type             string `json:"type"`
	} `json:"subject"`
	Reason     string    `json:"reason"`
	Unread     bool      `json:"unread"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastReadAt time.Time `json:"last_read_at"`
	URL        string    `json:"url"`
}

type Notifications []Notification

func (ns *Notifications) get(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "token "+config.GitHubToken)

	date, err := lm.Read()
	if err != nil {
		return err
	}

	if config.Polling && len(date) > 0 {
		logger.Info("Adding If-Modified-Since header")
		req.Header.Add("If-Modified-Since", string(date))
	}

	logger.Info("GET " + url)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	logger.Info("DONE " + resp.Status)

	if resp.StatusCode == 304 {
		logger.Info(resp.Status)
		return nil
	}

	if resp.StatusCode != 200 {
		return errors.New("Unable to access to the endpoint: " + url)
	}

	if resp.Header.Get("Last-Modified") != "" {
		date := []byte(resp.Header.Get("Last-Modified"))
		logger.Info("Last-Modified: " + string(date))
		lm.Write(date)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &ns)
	if err != nil {
		return err
	}

	return nil
}

func (n *Notification) markAsRead() error {
	req, err := http.NewRequest("PATCH", n.URL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "token "+config.GitHubToken)
	logger.Info("Marking the notification as read: " + n.URL)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	logger.Info("DONE " + resp.Status)

	if resp.StatusCode != 205 {
		return errors.New("Failed to mark the notification as read")
	}

	return nil
}
