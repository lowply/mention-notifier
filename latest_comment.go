package main

import (
	"encoding/json"
	"time"
)

type LatestComment struct {
	URL                 string `json:"url"`
	ID                  int    `json:"id"`
	PullRequestReviewID int    `json:"pull_request_review_id"`
	DiffHunk            string `json:"diff_hunk"`
	Path                string `json:"path"`
	Position            int    `json:"position"`
	OriginalPosition    int    `json:"original_position"`
	CommitID            string `json:"commit_id"`
	OriginalCommitID    string `json:"original_commit_id"`
	InReplyToID         int    `json:"in_reply_to_id"`
	User                struct {
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
	} `json:"user"`
	Body           string    `json:"body"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	HTMLURL        string    `json:"html_url"`
	PullRequestURL string    `json:"pull_request_url"`
	Links          struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		PullRequest struct {
			Href string `json:"href"`
		} `json:"pull_request"`
	} `json:"_links"`
}

func newLatestComment(url string) (*LatestComment, error) {
	_, bytes, err := getURL(url)
	if err != nil {
		return nil, err
	}

	var comment LatestComment
	err = json.Unmarshal(bytes, &comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
