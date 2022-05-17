# notifier

A golang package to send notification on some apps.

## Usage

### Mattermost

To set up a webhook, check the [mattermost documentation](https://docs.mattermost.com/developer/webhooks-incoming.html)

```go
url := "https://framateam.org/hooks/your-mattermost-hook-url"
n := mattermost.NewNotifier(url)
n.Notify("Hello, world!")
```
