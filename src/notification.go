package main

import (
	"log"
	"os"
	"strings"
)

type notification struct {
	Reason     string `json:"reason"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	Subject struct {
		Title            string `json:"title"`
		URL              string `json:"url"`
		LatestCommentURL string `json:"latest_comment_url"`
		Type             string `json:"type"`
	} `json:"subject"`
	URL string `json:"url"`
	c   comment
}

func (n *notification) check() (bool, error) {
	reason := "mention"
	if os.Getenv("MN_REASON") != "" {
		reason = os.Getenv("MN_REASON")
	}

	if n.Reason != reason {
		return true, nil
	}

	if n.Subject.Type == "RepositoryVulnerabilityAlert" {
		return true, nil
	}

	if n.Subject.LatestCommentURL == "" {
		log.Println("Empty LatestCommentURL: " + n.Subject.URL)
		return true, nil
	}

	if !strings.Contains(n.Subject.LatestCommentURL, "comments") {
		// If the latest comment URL is not a comment URL, these are the possible scenarios:
		//
		// 1. Issue is newly opened and mentioned
		// 2. Issue is closed
		// 3. Issue is closed and mentioned
		// 4. Issue is reopened
		//
		// We do the following:
		//
		// 1. There shouldn't be closed or reopened event, so just do nothing here
		// 2. Skip
		// 3. Will be catched ealier as the reason will be "mention"
		// 4. Skip
		//
		// Therefore, checking if the latest events has either "closed" or "reopened" should be enough
		log.Println("The latest comment URL is not a comment URL: " + n.Subject.LatestCommentURL)
		log.Println("Checking the events of the issue/pr...")
		ia := newIssueEventsAPI(n.Subject.URL, n.Subject.Type)
		is, err := ia.get()
		if err != nil {
			return true, err
		}
		for i := len(is) - 1; i >= 0; i-- {
			if is[i].closedOrReopened() {
				return true, nil
			}
		}
	}

	c := newCommentAPI(n.Subject.LatestCommentURL)
	comment, err := c.get()
	if err != nil {
		return true, err
	}
	n.c = comment

	login := os.Getenv("GITHUB_ACTOR")
	if !n.c.mentioned(login) {
		return true, nil
	}

	return false, nil
}

func (n *notification) notify() error {
	log.Println("Posting to Slack...")
	s := newSlackAPI()
	err := s.post(n)
	if err != nil {
		return err
	}
	return nil
}

func (n *notification) markAsRead() error {
	log.Println("Marking the thread read...")
	q := newQuery()
	err := q.patch(n.URL)
	if err != nil {
		return err
	}
	return nil
}
