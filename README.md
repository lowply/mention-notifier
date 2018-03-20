# Mention Notifier

Get notified on Slack when mentioned on GitHub.

## Getting started

1. Setup [Incoming Webhooks](https://api.slack.com/incoming-webhooks) on your Slack
1. Get a [Personal access tokens](https://github.com/settings/tokens) with `notifications` and `repo` scopes
1. Setup your AWS credential with appropriate permission
1. `mv cf_example.yml cf.yml` and update it
1. Run `AWSPROFILE=profile ./cf deploy`

## Development

You can run this locally by doing following:

```bash
LOGIN=login \
GITHUB_TOKEN=token \
SLACK_ENDPOINT=endpoint \
GITHUB_ENDPOINT=https://api.github.com/notifications \
REASON=mention \
POLLING=false \
make run
```

The `POLLING` option should be `false` while development.
