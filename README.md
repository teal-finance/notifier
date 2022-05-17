# notifier

![workflow](https://github.com/teal-finance/notifier/actions/workflows/go.yml/badge.svg)
[![Go Report](https://goreportcard.com/badge/github.com/teal-finance/notifier)](https://goreportcard.com/report/github.com/teal-finance/notifier)
[![Go Version](https://img.shields.io/github/go-mod/go-version/kahlys/alfred.svg)](https://github.com/teal-finance/notifier)
[![Go Reference](https://pkg.go.dev/badge/github.com/teal-finance/notifier.svg)](https://pkg.go.dev/github.com/teal-finance/notifier)

A golang package to send notification on some apps.

## Usage

### Mattermost

To set up a webhook, check the [mattermost documentation](https://docs.mattermost.com/developer/webhooks-incoming.html)

```go
url := "https://framateam.org/hooks/your-mattermost-hook-url"
n := mattermost.NewNotifier(url)
n.Notify("Hello, world!")
```
