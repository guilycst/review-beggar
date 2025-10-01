package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SlackMessage struct {
	Text string `json:"text"`
}

func main() {
	// Get inputs from environment variables
	slackWebhookURL := os.Getenv("INPUT_SLACK_WEBHOOK_URL")
	if slackWebhookURL == "" {
		fmt.Println("Error: SLACK_WEBHOOK_URL input is required")
		os.Exit(1)
	}

	// Get GitHub context from environment
	prNumber := os.Getenv("GITHUB_PR_NUMBER")
	prTitle := os.Getenv("GITHUB_PR_TITLE")
	prURL := os.Getenv("GITHUB_PR_URL")
	prAuthor := os.Getenv("GITHUB_PR_AUTHOR")
	repoName := os.Getenv("GITHUB_REPOSITORY")

	// Validate required inputs
	if prNumber == "" || prTitle == "" || prURL == "" {
		fmt.Println("Error: Missing required PR information")
		os.Exit(1)
	}

	// Format Slack message
	message := formatSlackMessage(prNumber, prTitle, prURL, prAuthor, repoName)

	// Send message to Slack
	if err := sendSlackMessage(slackWebhookURL, message); err != nil {
		fmt.Printf("Error sending Slack message: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent Slack notification")
}

func formatSlackMessage(prNumber, prTitle, prURL, prAuthor, repoName string) string {
	author := prAuthor
	if author == "" {
		author = "Unknown"
	}

	repo := repoName
	if repo == "" {
		repo = "Repository"
	}

	return fmt.Sprintf(
		":eyes: *Pull Request Ready for Review* :eyes:\n\n"+
			"*Repository:* %s\n"+
			"*PR #%s:* %s\n"+
			"*Author:* %s\n"+
			"*Link:* <%s|View Pull Request>",
		repo, prNumber, prTitle, author, prURL,
	)
}

func sendSlackMessage(webhookURL, message string) error {
	slackMsg := SlackMessage{Text: message}

	payload, err := json.Marshal(slackMsg)
	if err != nil {
		return fmt.Errorf("failed to marshal Slack message: %w", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Slack API returned non-OK status: %d", resp.StatusCode)
	}

	return nil
}
