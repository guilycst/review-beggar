#!/bin/sh

set -e

# Extract PR information from GitHub context
export GITHUB_PR_NUMBER=$(jq -r '.pull_request.number // empty' "$GITHUB_EVENT_PATH")
export GITHUB_PR_TITLE=$(jq -r '.pull_request.title // empty' "$GITHUB_EVENT_PATH")
export GITHUB_PR_URL=$(jq -r '.pull_request.html_url // empty' "$GITHUB_EVENT_PATH")
export GITHUB_PR_AUTHOR=$(jq -r '.pull_request.user.login // empty' "$GITHUB_EVENT_PATH")

# Run the Go binary
/review-beggar
