package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	repo := os.Getenv("REPO") // e.g., platformoon/platformoon-core
	sha := os.Getenv("SHA")
	url := fmt.Sprintf("https://api.github.com/repos/%s/statuses/%s", repo, sha)
	state := os.Getenv("STATE")
	targetUrl := os.Getenv("TARGET_URL")
	description := os.Getenv("DESCRIPTION")

	entries, err := os.ReadDir("/etc/gh-checkmoon")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}

	token, err := os.ReadFile("/etc/gh-checkmoon/token")

	if err != nil {
		panic(any(err))
	}

	log.Println(string(token))

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
		panic(any(err))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	log.Println(string(body))
}
