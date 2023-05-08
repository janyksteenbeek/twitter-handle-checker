package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const twitterBaseURL = "https://twitter.com/"

func main() {
	username := os.Getenv("CHECKER_USERNAME")
	pushoverToken := os.Getenv("CHECKER_PUSHOVER_TOKEN")

	if username == "" || pushoverToken == "" {
		fmt.Println("Missing environment variables.")
		return
	}

	for {
		checkTwitterHandle(username, pushoverToken)
		time.Sleep(1 * time.Hour)
	}
}

func checkTwitterHandle(username, pushoverToken string) {
	fmt.Printf("Checking Twitter handle @%s...\n", username)

	resp, err := http.Get(twitterBaseURL + username)

	if err != nil {
		fmt.Printf("Error getting Twitter handle: %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		sendPushoverNotification(username, pushoverToken)
	}
}

func sendPushoverNotification(username, pushoverToken string) {
	msg := fmt.Sprintf("Twitter handle @%s is now available!", username)
	data := fmt.Sprintf("token=%s&user=%s&message=%s", pushoverToken, username, msg)

	resp, err := http.Post("https://api.pushover.net/1/messages.json", "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Printf("Error sending Pushover notification: %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Pushover notification sent successfully!")
	} else {
		fmt.Printf("Error sending Pushover notification: %s\n", resp.Status)
	}
}
