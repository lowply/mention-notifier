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

	var ns Notifications
	err = ns.Get(config.GitHubEndpoint)
	if err != nil {
		return err
	}

	if len(ns) == 0 {
		logger.Info("No notificaions")
		return nil
	}

	for _, n := range ns {
		if n.Reason != config.Reason {
			continue
		}

		if n.Subject.LatestCommentURL == "" {
			logger.Info("Empty LatestCommentURL: " + n.Subject.URL)
			continue
		}

		if !strings.Contains(n.Subject.LatestCommentURL, "comments") {
			logger.Info("Issue or Pull Request was closed or reopened: " + n.Subject.LatestCommentURL)
			continue
		}

		var c LatestComment
		err := c.Get(n.Subject.LatestCommentURL)
		if err != nil {
			return err
		}

		if !strings.Contains(c.Body, config.Login) {
			logger.Info("There is a notification, but the latest comment didn't mention you.")
			continue
		}

		var s = Slack{
			Notification: n,
			Comment:      c,
		}

		err = s.Post()
		if err != nil {
			return err
		}

		err = n.MarkAsRead()
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