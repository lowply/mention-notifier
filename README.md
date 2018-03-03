# Mention Notifier

Get notified on Slack when mentioned on GitHub.

## Getting started

1. Setup [Incoming Webhooks](https://api.slack.com/incoming-webhooks) on your Slack
1. Get a [Personal access tokens](https://github.com/settings/tokens) with `notifications` and `repo` scopes
1. Download the *mention-notifier_linux_amd64.zip*
1. Unzip it and move it to wherever under your $PATH
1. Put the *mention-notifier.json* file in the `~/.config` dir
1. Add it to crontab

#### Config example

Use the *mention-notifier.json.example* file as your template.

```json
{
    "Login": "login",
    "GitHubToken": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "GitHubEndpoint": "https://api.github.com/notifications",
    "SlackEndpoint": "https://hooks.slack.com/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "Reason": "mention",
    "Polling": true
}
```

#### Config description

- `Login`: Your GitHub username
  - Required / Default value: `nil`
- `GitHubToken`: Your personal access token for GitHub
  - Required / Default value: `nil`
- `GitHubEndpoint`: GitHub's Notification API endpoint
  - Required / Default value: `https://api.github.com/notifications`
- `SlackEndpoint`: Slack's Incoming Webhooks endpoint URL.
  - Required / Default value: `nil`
- `Reason`: Reason for notification. See [Notification Reasons](https://developer.github.com/v3/activity/notifications/#notification-reasons) for other options.
  - Optional / Default value: `mention`
- `Polling`: Enables the "Last-Modified" header checking. See [Notifications](https://developer.github.com/v3/activity/notifications/) for more details.
  - Optional / Default value: `true`

#### Crontab exemple

```
* * * * * /home/sho/bin/mention-notifier >/dev/null
```
