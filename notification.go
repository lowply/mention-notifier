package main

import "time"

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
