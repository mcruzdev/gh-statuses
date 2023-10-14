package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	repo := os.Getenv("REPO") // e.g., platformoon/platformoon-core
	// _ := os.Getenv("SHA")
	url := fmt.Sprintf("https://api.github.com/repos/%s/statuses/%s", repo, "abcd5c018e5595e89d0283d783bdb10a2b66dca5")
	state := os.Getenv("STATE")
	targetUrl := os.Getenv("TARGET_URL")
	description := os.Getenv("DESCRIPTION")
	token, err := os.ReadFile("/etc/gh-token")

	requestBody := []byte(fmt.Sprintf(`{
		"state": "%s",
		"target_url": "%s",
		"description": "%s",
		"context": "continuous-integration/platformoon"
	}`, state, targetUrl, description))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(fmt.Errorf("error while creating request"))
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+string(token))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Errorf("error while requesting Github API"))
	}
	defer resp.Body.Close()
	log.Println("Commit status setted")
}
