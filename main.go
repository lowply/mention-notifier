package main

import (
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
		return nil
	}

	for _, n := range ns {
		if n.Reason != config.Reason {
			continue
		}

		if n.Subject.LatestCommentURL == "" {
			logger.Info("Empty LatestCommentURL for: " + n.Subject.URL)
			continue
		}

		var c LatestComment
		err := c.Get(n.Subject.LatestCommentURL)
		if err != nil {
			return err
		}

		var s = Slack{
			Notification: n,
			Comment:      c,
		}

		err = s.Post()
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
