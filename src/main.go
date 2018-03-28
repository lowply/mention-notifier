package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() error {
	err := config.Read()
	if err != nil {
		return err
	}

	var ns = new(Notifications)
	err = ns.query(config.GitHubEndpoint)
	if err != nil {
		return err
	}

	if len(*ns) == 0 {
		logger.Info("No notificaions")
		return nil
	}

	for _, n := range *ns {
		if n.Reason != config.Reason {
			continue
		}

		if n.Subject.LatestCommentURL == "" {
			logger.Info("Empty LatestCommentURL: " + n.Subject.URL)
			continue
		}

		if !strings.Contains(n.Subject.LatestCommentURL, "comments") {
			logger.Info("The latest comment URL is the issue URL: " + n.Subject.URL)
			logger.Info("Checking the events of the issue/pr...")

			var es = new(IssueEvents)
			err := es.query(n.Subject.URL + "/events")
			if err != nil {
				return err
			}

			if es.closedOrReopened() {
				logger.Info("Skipping notification as the issue is closed or reopened.")
				continue
			}
		}

		var c = new(LatestComment)
		err := c.query(n.Subject.LatestCommentURL)
		if err != nil {
			return err
		}

		if !strings.Contains(c.Body, config.Login) {
			logger.Info("There is a notification, but the latest comment didn't mention you.")
			continue
		}

		var s = new(Slack)
		s.Notification = n
		s.Comment = c

		err = s.post()
		if err != nil {
			return err
		}

		err = n.markAsRead()
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if os.Getenv("LOCAL") == "true" {
		err := handler()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		lambda.Start(handler)
	}
}
