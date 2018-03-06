package main

import (
	"log"
	"os"
)

func main() {
	err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	var ns Notifications
	err = ns.Get(config.GitHubEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	if len(ns) == 0 {
		os.Exit(0)
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
			log.Fatal(err)
		}

		var s = Slack{
			Notification: n,
			Comment:      c,
		}

		err = s.Post()
		if err != nil {
			log.Fatal(err)
		}
	}
}
