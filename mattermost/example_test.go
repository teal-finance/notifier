package mattermost_test

import "github.com/teal-finance/notifier/mattermost"

func ExampleNotifier_Notify() {
	url := "https://framateam.org/hooks/your-mattermost-hook-url"
	n := mattermost.NewNotifier(url)
	n.Notify("Hello, world!")
}
