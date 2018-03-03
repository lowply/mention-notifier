package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Slack struct {
	Endpoint     string
	Notification Notification
}

func (s *Slack) Post() error {
	fmt.Println("Posting to Slack...")

	n := s.Notification
	comment, err := newLatestComment(n.Subject.LatestCommentURL)
	if err != nil {
		return err
	}

	pi := strconv.FormatInt(comment.UpdatedAt.Unix(), 10)
	data := url.Values{}
	data.Set("payload", `{
		"channel": "#notifications",
		"username": "GitHub Notifier",
		"icon_emoji": ":octocat:",
		"attachments": [
			{
				"fallback": "`+n.Subject.Title+`",
				"color": "#36a64f",
				"pretext": "Hey @lowply, you've got a new mention!",
				"author_name": "`+comment.User.Login+`",
				"author_link": "`+comment.User.HTMLURL+`",
				"author_icon": "`+comment.User.AvatarURL+`",
				"title": "`+n.Subject.Title+`",
				"title_link": "`+comment.HTMLURL+`",
				"text": "Repository: `+n.Repository.FullName+`\n`+comment.Body+`",
				"ts": "`+pi+`"
			}
		]
	}`)
	req, err := http.NewRequest("POST", s.Endpoint, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	return nil
}
