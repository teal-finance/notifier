package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Notifier for sending messages to a Mattermost server.
type notifierMattermost struct {
	endpoint string
}

// NewNotifier creates a new Notifier given a Mattermost server endpoint (see mattermost hooks).
func WithMattermost(endpoint string) Notifier {
	return notifierMattermost{endpoint}
}

type mattermostRequest struct {
	Text string `json:"text,omitempty"`
}

// Notify sends a message to the Mattermost server.
func (n notifierMattermost) Notify(msg string) error {
	param := mattermostRequest{Text: msg}
	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(param)
	if err != nil {
		return fmt.Errorf("notifier mattermost: %w", err)
	}

	resp, err := http.Post(n.endpoint, "application/json", reqBody)
	if err != nil {
		return fmt.Errorf("notifier mattermost: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("notifier mattermost: http code status %v", resp.Status)
	}
	return nil
}
