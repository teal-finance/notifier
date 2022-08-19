package mattermost

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

// Notifier for sending messages to a Mattermost server.
type Notifier struct {
	endpoint string
}

// NewNotifier creates a new Notifier given a Mattermost server endpoint (see mattermost hooks).
func NewNotifier(endpoint string) Notifier {
	return Notifier{endpoint}
}

// Notify sends a message to the Mattermost server.
func (n Notifier) Notify(msg string) error {
	resp, err := http.Post(
		n.endpoint,
		"application/json",
		bytes.NewBuffer([]byte(`{"text":"`+msg+`"}`)),
	)
	if err != nil {
		return fmt.Errorf("mattermost notify: %w from host=%s", err, n.host())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("mattermost notify: %s from host=%s", resp.Status, n.host())
	}
	return nil
}

func (n Notifier) host() string {
	if url, er := url.Parse(n.endpoint); er == nil {
		return url.Hostname()
	}
	return ""
}
