package main

import (
	"log"
	"os"
)

func main() {
	// Required
	required := []string{"GITHUB_ACTOR", "_GITHUB_TOKEN", "SLACK_ENDPOINT"}
	for _, v := range required {
		if os.Getenv(v) == "" {
			log.Fatal(v + " is empty.")
		}
	}

	na := newNotificationAPI()
	ns, err := na.get()
	if err != nil {
		log.Fatal(err)
	}

	if len(ns) == 0 {
		log.Println("No notifications.")
		return
	}

	for _, n := range ns {
		skip, err := n.check()
		if err != nil {
			log.Fatal(err)
		}
		if skip {
			continue
		}

		err = n.notify()
		if err != nil {
			log.Fatal(err)
		}

		err = n.markAsRead()
		if err != nil {
			log.Fatal(err)
		}
	}
}
