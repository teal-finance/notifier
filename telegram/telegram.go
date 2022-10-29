// copied from ftxnft notifier
package telegram

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

// Bot is a telegram bot for a specific chat room
type Bot struct {
	botToken string
	chatID   string
	clt      *http.Client
}

func NewTelegramBotClient(botToken, chatID string) Bot {
	return Bot{
		botToken: botToken,
		chatID:   chatID,
		clt:      &http.Client{},
	}
}

func (b Bot) Notify(text string) error {
	var telegramApi string = "https://api.telegram.org/bot" + b.botToken + "/sendMessage"

	response, err := b.clt.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {b.chatID},
			"text":    {text},
		})
	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return err
	}
	defer response.Body.Close()

	resp := sendMessageResponseParam{}
	if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
		log.Printf("error in parsing telegram answer %s", err.Error())
		return err
	}
	log.Printf("Sending notification via telegram: %v", resp.Ok)

	return nil
}

type sendMessageResponseParam struct {
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
