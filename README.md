# Mention Notifier

A GitHub Action that notifies you on Slack when you're mentioned on GitHub.

## Environment variables

- `_GITHUB_TOKEN` (*required*): Your personal access token with [the `notification` scope](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line). The `GITHUB_TOKEN` env var that comes with the workflow by default doesn't have the `notification` scope, so you'll need this env var instead. Remember to [Authorize](https://help.github.com/en/articles/authorizing-a-personal-access-token-for-use-with-a-saml-single-sign-on-organization) it if you're a part of an organization that uses [SAML single sign-on](https://help.github.com/en/articles/about-authentication-with-saml-single-sign-on).
- `SLACK_ENDPOINT` (*required*): Slack API endpoint URL
- `MN_INTERVAL` (*optional*): The interval that used to add the If-Modified-Since HTTP header. Change this value when you changed the workflow interval. Default is `1`. Should be in the range of `1-59`.

## Workflow example

```
workflow "Run Mention Notifier" {
  on = "schedule(* * * * *)"
  resolves = ["Mention Notifier"]
}

action "Mention Notifier" {
  uses = "lowply/mention-notifier@0.0.1"
  secrets = ["SLACK_ENDPOINT", "_GITHUB_TOKEN"]
}
```

## Development

You can run this locally by running following command:

```bash
export _GITHUB_TOKEN=token
export SLACK_ENDPOINT=endpoint
export GITHUB_ACTOR=login
export MN_POLLING=false
cd src
go run .
```

Note that the `MN_POLLING` option is recommended to be set as `false` while development.
