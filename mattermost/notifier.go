package mattermost

import (
	"bytes"
	"fmt"
	"net/http"
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
		return fmt.Errorf("notify: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("notify: http status code %d", resp.StatusCode)
	}
	return nil
}
