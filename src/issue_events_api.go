package main

import (
	"encoding/json"
	"log"
	"strings"
)

type issueEventsAPI struct {
	*query
	endpoint string
}

func newIssueEventsAPI(url, typ string) *issueEventsAPI {
	q := newQuery()
	i := &issueEventsAPI{query: q}

	// The issue event endpoint is fixed to /repos/:owner/:repo/issues.
	// If the notification's subject.url value is a URL ends with /pulls/:id
	// we need to replace it to /issues/:id.
	// Ref: https://developer.github.com/v3/issues/events/#get-a-single-event
	if typ == "PullRequest" {
		url = strings.Replace(url, "/pulls/", "/issues/", 1)
	}

	i.endpoint = url + "/events"
	return i
}

func (ia *issueEventsAPI) parseLink(link string) map[string]string {
	links := map[string]string{}
	for _, l := range strings.Split(link, ",") {
		l = strings.TrimSpace(l)
		m := strings.Split(l, ";")

		url := strings.TrimSpace(m[0])
		url = strings.TrimPrefix(url, "<")
		url = strings.TrimSuffix(url, ">")

		rel := strings.TrimSpace(m[1])
		rel = strings.TrimPrefix(rel, "rel=")
		rel = strings.Trim(rel, "\"")

		links[rel] = url
	}
	return links
}

func (ia *issueEventsAPI) getTheLastURL() error {
	link, err := ia.query.link(ia.endpoint)
	if err != nil {
		return err
	}
	if link != "" {
		links := ia.parseLink(link)
		ia.endpoint = links["last"]
	}
	return nil
}

func (ia *issueEventsAPI) get() ([]issueEvent, error) {
	log.Println("Getting the last URL in the Link header...")
	err := ia.getTheLastURL()
	if err != nil {
		return nil, err
	}
	log.Println("Getting the last events...")
	data, err := ia.query.get(ia.endpoint)
	if err != nil {
		return nil, err
	}

	is := []issueEvent{}
	err = json.Unmarshal(data, &is)
	if err != nil {
		return nil, err
	}

	return is, nil
}
