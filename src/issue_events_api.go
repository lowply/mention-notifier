package main

import (
	"encoding/json"
	"os"
	"strings"
)

type issueEventsAPI struct {
	*query
	endpoint      string
	endpoint_last string
	token         string
}

func newIssueEventsAPI(url string) *issueEventsAPI {
	q := newQuery()
	i := &issueEventsAPI{query: q}
	i.token = os.Getenv("GITHUB_TOKEN")
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
	logger.Info("Getting the last URL in the Link header...")
	err := ia.getTheLastURL()
	if err != nil {
		return nil, err
	}
	logger.Info("Getting the last events...")
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
