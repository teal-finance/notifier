package notifier

import (
	"os"
)

func ExampleWithWriter() {
	n := WithWriter(os.Stdout)
	n.Notify("hello")
	// Output: hello
}

func ExampleWithMattermost() {
	url := "https://framateam.org/hooks/your-mattermost-hook-url"
	n := WithMattermost(url)
	n.Notify("hello")
}

func ExampleWithTelegram() {
	url := "https://framateam.org/hooks/your-mattermost-hook-url"
	n := WithMattermost(url)
	n.Notify("hello")
}
