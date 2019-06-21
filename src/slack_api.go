package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type slackAPI struct {
	endpoint string
	login    string
}

func newSlackAPI() *slackAPI {
	s := &slackAPI{}
	s.endpoint = os.Getenv("SLACK_ENDPOINT")
	s.login = os.Getenv("GITHUB_ACTOR")
	return s
}

func (s *slackAPI) post(n *notification) error {
	updatedAt := strconv.FormatInt(n.c.UpdatedAt.Unix(), 10)
	data := url.Values{}
	data.Set("payload", `{
		"username": "Mention Notifier",
		"icon_emoji": ":octocat:",
		"attachments": [
			{
				"fallback": "`+n.Subject.Title+`",
				"color": "#36a64f",
				"pretext": "Hey @`+s.login+`, you've got a new mention!",
				"author_name": "`+n.c.User.Login+`",
				"author_link": "`+n.c.User.HTMLURL+`",
				"author_icon": "`+n.c.User.AvatarURL+`",
				"title": "`+n.Subject.Title+`",
				"title_link": "`+n.c.HTMLURL+`",
				"text": "Repository: `+n.Repository.FullName+`\n`+n.c.Body+`",
				"ts": "`+updatedAt+`"
			}
		]
	}`)
	req, err := http.NewRequest("POST", s.endpoint, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Error posting to Slack: " + resp.Status)
	}

	defer resp.Body.Close()
	log.Println("DONE " + resp.Status)

	return nil
}
