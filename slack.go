package main

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Slack struct {
	Notification Notification
	Comment      LatestComment
}

func (s *Slack) Post() error {
	updated_at := strconv.FormatInt(s.Comment.UpdatedAt.Unix(), 10)
	data := url.Values{}
	data.Set("payload", `{
		"channel": "#notifications",
		"username": "GitHub Notifier",
		"icon_emoji": ":octocat:",
		"attachments": [
			{
				"fallback": "`+s.Notification.Subject.Title+`",
				"color": "#36a64f",
				"pretext": "Hey @lowply, you've got a new mention!",
				"author_name": "`+s.Comment.User.Login+`",
				"author_link": "`+s.Comment.User.HTMLURL+`",
				"author_icon": "`+s.Comment.User.AvatarURL+`",
				"title": "`+s.Notification.Subject.Title+`",
				"title_link": "`+s.Comment.HTMLURL+`",
				"text": "Repository: `+s.Notification.Repository.FullName+`\n`+s.Comment.Body+`",
				"ts": "`+updated_at+`"
			}
		]
	}`)
	req, err := http.NewRequest("POST", config.SlackEndpoint, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	logger.Info("Posting to Slack")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	logger.Info("DONE " + resp.Status)

	return nil
}
