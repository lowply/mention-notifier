package main

import (
	"encoding/json"
	"os"
)

type commentAPI struct {
	*query
	login     string
	endpoint  string
	mentioned bool
}

func newCommentAPI(url string) *commentAPI {
	q := newQuery()
	c := &commentAPI{query: q}
	c.login = os.Getenv("GITHUB_ACTOR")
	c.endpoint = url
	c.mentioned = false
	return c
}

func (c *commentAPI) get() (comment, error) {
	comment := &comment{}
	data, err := c.query.get(c.endpoint)
	if err != nil {
		return *comment, err
	}
	err = json.Unmarshal(data, &comment)
	if err != nil {
		return *comment, err
	}
	return *comment, nil
}
