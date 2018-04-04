package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response interface {
}

func Query(url string, res Response) error {
	checkLastModified := false
	if strings.Contains(url, "/notifications") {
		checkLastModified = true
	}
	httpreq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	httpreq.Header.Add("Authorization", "token "+config.GitHubToken)

	if checkLastModified {
		date, err := lm.Read()
		if err != nil {
			return err
		}

		if config.Polling && len(date) > 0 {
			logger.Info("Adding If-Modified-Since header")
			httpreq.Header.Add("If-Modified-Since", string(date))
		}
	}

	logger.Info("GET " + url)
	client := http.Client{}
	httpres, err := client.Do(httpreq)
	if err != nil {
		return err
	}
	defer httpres.Body.Close()
	logger.Info("DONE " + httpres.Status)

	if checkLastModified && httpres.StatusCode == 304 {
		return nil
	}

	if httpres.StatusCode != 200 {
		return errors.New("Unable to access to the endpoint: " + url)
	}

	if checkLastModified && httpres.Header.Get("Last-Modified") != "" {
		date := []byte(httpres.Header.Get("Last-Modified"))
		logger.Info("Last-Modified: " + string(date))
		lm.Write(date)
	}

	bytes, err := ioutil.ReadAll(httpres.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return err
	}

	return nil
}
