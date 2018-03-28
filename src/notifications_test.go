package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotificationsGet(t *testing.T) {
	mock, err := ioutil.ReadFile("mock/notifications.json")
	if err != nil {
		log.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(mock))
	}))
	defer ts.Close()

	var ns = new(Notifications)
	err = ns.get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	if (*ns)[0].ID != "318027093" {
		t.Fatalf("Got : %v, Expected : %v", (*ns)[0].ID, "318027093")
	}

	if (*ns)[0].Reason != "mention" {
		t.Fatalf("Got : %v, Expected : %v", (*ns)[0].Reason, "mention")
	}

	if !(*ns)[1].Unread {
		t.Fatalf("Got : %v, Expected : %v", (*ns)[1].Unread, true)
	}
}
