package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getURL(url string) (int, []byte, error) {
	if url == "" {
		return 0, nil, errors.New("URL can't be nil")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, nil, err
	}

	req.Header.Add("Authorization", "token "+config.GitHubToken)

	if url == config.GitHubEndpoint {
		date, err := lm.Read()
		if err != nil {
			return 0, nil, err
		}

		if config.Polling && len(date) > 0 {
			req.Header.Add("If-Modified-Since", string(date))
		}
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	if resp.Header.Get("Last-Modified") != "" {
		date := []byte(resp.Header.Get("Last-Modified"))
		fmt.Println(string(date))
		lm.Write(date)
	}

	return resp.StatusCode, bytes, nil
}

func main() {
	err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	code, bytes, err := getURL(config.GitHubEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	if code == 200 {
		var notifications []Notification
		err = json.Unmarshal(bytes, &notifications)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range notifications {
			if v.Reason == config.Reason && v.Subject.LatestCommentURL != "" {
				var s = Slack{
					Endpoint:     config.SlackEndpoint,
					Notification: v,
				}
				err := s.Post()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	} else {
		fmt.Println(code)
	}
}
