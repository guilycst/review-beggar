# review-beggar

Sends Slack messages to a channel when a Pull Request is ready to be reviewed.

## Usage

Add this action to your workflow to notify a Slack channel when a PR is ready for review:

```yaml
name: PR Review Notification
on:
  pull_request:
    types: [opened, ready_for_review]

jobs:
  notify-slack:
    runs-on: ubuntu-latest
    steps:
      - name: Notify Slack
        uses: guilycst/review-beggar@main
        with:
          slack_webhook_url: ${{ secrets.SLACK_WEBHOOK_URL }}
```

## Inputs

| Input | Description | Required |
|-------|-------------|----------|
| `slack_webhook_url` | Slack webhook URL for sending notifications | Yes |

## Setup

1. Create a Slack webhook URL:
   - Go to https://api.slack.com/messaging/webhooks
   - Create a new webhook for your workspace
   - Copy the webhook URL

2. Add the webhook URL as a secret in your GitHub repository:
   - Go to your repository settings
   - Navigate to Secrets and variables > Actions
   - Add a new secret named `SLACK_WEBHOOK_URL`
   - Paste your Slack webhook URL as the value

3. Add the workflow file to your repository as shown in the usage example

## Features

- Sends formatted Slack notifications when PRs are opened or marked as ready for review
- Includes PR number, title, author, and direct link to the PR
- Simple setup with minimal configuration required

## License

MIT License - see [LICENSE](LICENSE) for details
