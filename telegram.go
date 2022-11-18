package notifier

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Bot is a telegram bot for a specific chat room
type notifierTelegram struct {
	chatID   string
	endpoint string
}

func WithTelegram(botToken, chatID string) Notifier {
	return notifierTelegram{
		chatID:   chatID,
		endpoint: fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", botToken),
	}
}

type telegramResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int    `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID    int64  `json:"id"`
			Title string `json:"title"`
			Type  string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

func (n notifierTelegram) Notify(msg string) error {
	response, err := http.PostForm(
		n.endpoint,
		url.Values{
			"chat_id": {n.chatID},
			"text":    {msg},
		})
	if err != nil {
		return fmt.Errorf("notifier telegram: %w", err)
	}
	defer response.Body.Close()

	resp := telegramResponse{}
	if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return fmt.Errorf("notifier telegram: %w", err)
	}

	if !resp.Ok {
		return fmt.Errorf("notifier telegram: sending request failed")
	}

	return nil
}
