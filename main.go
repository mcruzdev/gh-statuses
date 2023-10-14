package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	repo := os.Getenv("REPO") // e.g., platformoon/platformoon-core
	sha := os.Getenv("SHA")
	url := fmt.Sprintf("https://api.github.com/repos/%s/statuses/%s", repo, sha)
	state := os.Getenv("STATE")
	targetUrl := os.Getenv("TARGET_URL")
	description := os.Getenv("DESCRIPTION")

	token, err := os.ReadFile("/etc/gh-checkmoon/token")

	if err != nil {
		panic(any(err))
	}

	requestBody := []byte(fmt.Sprintf(`{
		"state": "%s",
		"target_url": "%s",
		"description": "%s",
		"context": "build"
	}`, state, targetUrl, description))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(any(err))
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", strings.TrimSpace(string(token))))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(any(err))
	}
	defer resp.Body.Close()
}
