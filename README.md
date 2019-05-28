# Mention Notifier

A GitHub Action that notifies you on Slack when you're mentioned on GitHub.

## Environment variables

- `SLACK_ENDPOINT` (*required*): Slack API endpoint URL
- `MN_INTERVAL` (*optional*): The interval that used to add the If-Modified-Since HTTP header. Change this value when you changed the workflow interval. Default is `1`. Should be in the range of `1-59`.

## Workflow example

```
workflow "Mention Notifier" {
  resolves = ["Run Mention Notifier"]
  on = "schedule(* * * * *)"
}

action "Run Mention Notifier" {
  uses = "lowply/mention-notifier@0.0.8"
  secrets = {
    "SLACK_ENDPOINT",
    "GITHUB_TOKEN"
}

```

## Development

You can run this locally by running following command:

```bash
export GITHUB_ACTOR=login
export GITHUB_TOKEN=token
export SLACK_ENDPOINT=endpoint
export MN_POLLING=false
cd src
go run .
```

The `MN_POLLING` option should be set to `false` while development.
