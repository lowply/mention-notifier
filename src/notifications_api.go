package main

import (
	"encoding/json"
	"os"
)

type notificationAPI struct {
	*query
	endpoint string
}

func newNotificationAPI() *notificationAPI {
	q := newQuery()
	na := &notificationAPI{query: q}

	na.endpoint = "https://api.github.com/notifications"
	if os.Getenv("GITHUB_ENDPOINT") != "" {
		na.endpoint = os.Getenv("GITHUB_ENDPOINT")
	}

	return na
}

func (na *notificationAPI) get() ([]notification, error) {
	data, err := na.query.get(na.endpoint + "?participating=true")
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, err
	}

	ns := &[]notification{}
	err = json.Unmarshal(data, &ns)
	if err != nil {
		return nil, err
	}

	return *ns, nil
}
