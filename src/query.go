package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type query struct {
	token    string
	polling  bool
	interval time.Duration
	lastRun  time.Time
}

func newQuery() *query {
	q := &query{}
	q.token = os.Getenv("_GITHUB_TOKEN")
	q.polling = true
	return q
}

func (q *query) parseEnv() error {
	if os.Getenv("MN_POLLING") != "" {
		b, err := strconv.ParseBool(os.Getenv("MN_POLLING"))
		if err != nil {
			return err
		}
		q.polling = b
	}

	i := ""
	if os.Getenv("MN_INTERVAL") != "" {
		i = os.Getenv("MN_INTERVAL")
	} else {
		i = "1"
	}

	interval, err := time.ParseDuration(i + "m")
	if err != nil {
		return err
	}

	q.interval = interval
	return nil
}

func (q *query) formatTime() string {
	f := "Mon, 2 Jan 2006 15:04:05 GMT"
	return q.lastRun.Format(f)
}

func (q *query) get(url string) ([]byte, error) {
	err := q.parseEnv()
	if err != nil {
		return nil, err
	}

	isNotificationAPI := strings.Contains(url, "/notifications")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+q.token)

	if q.polling && isNotificationAPI {
		// Adding the If-Modified-Since header to check if
		// there are new notifications since the last workflow run.
		// See https://developer.github.com/v3/activity/notifications/ for details.
		q.lastRun = time.Now().UTC().Add(-q.interval)
		logger.Info("Adding If-Modified-Since header")
		req.Header.Add("If-Modified-Since", q.formatTime())
	}

	logger.Info("GET " + url)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	logger.Info("DONE " + res.Status)

	if isNotificationAPI {
		logger.Info("X-RateLimit-Limit: " + res.Header.Get("X-RateLimit-Limit"))
		logger.Info("X-RateLimit-Remaining: " + res.Header.Get("X-RateLimit-Remaining"))
		logger.Info("X-RateLimit-Reset: " + res.Header.Get("X-RateLimit-Reset"))
	}

	if q.polling && res.StatusCode == 304 {
		logger.Info("304 Not Modified since: " + q.formatTime())
		return nil, nil
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Unable to access to the endpoint: " + url)
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (q *query) link(url string) (string, error) {
	link := ""
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return link, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+q.token)

	logger.Info("GET " + url)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return link, err
	}
	defer res.Body.Close()

	logger.Info("DONE " + res.Status)

	if res.StatusCode != 200 {
		return link, errors.New("Unable to access to the endpoint: " + url)
	}

	link = res.Header.Get("Link")

	return link, nil
}

func (q *query) patch(url string) error {
	req, err := http.NewRequest(http.MethodPatch, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "token "+q.token)
	logger.Info("Marking the notification as read: " + url)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	logger.Info("DONE " + resp.Status)

	if resp.StatusCode != 205 {
		// https://developer.github.com/v3/activity/notifications/#mark-a-thread-as-read
		return errors.New("Failed to mark the notification thread as read")
	}
	return nil
}
