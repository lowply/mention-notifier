package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIssueEventsGet(t *testing.T) {
	mock, err := ioutil.ReadFile("mock/issue_events.json")
	if err != nil {
		log.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(mock))
	}))
	defer ts.Close()

	var es = new(IssueEvents)
	err = es.get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	if (*es)[0].ID != 1539017820 {
		t.Fatalf("Got : %v, Expected : %v", (*es)[0].ID, 1539017820)
	}

	if (*es)[0].Actor.Login != "lowply" {
		t.Fatalf("Got : %v, Expected : %v", (*es)[0].Actor.Login, "lowply")
	}

	if (*es)[1].Event != "subscribed" {
		t.Fatalf("Got : %v, Expected : %v", (*es)[1].Event, "subscribed")
	}
}
