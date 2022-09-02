package notifier

import (
	"github.com/teal-finance/emo"
	"github.com/teal-finance/notifier/logger"
	"github.com/teal-finance/notifier/mattermost"
)

var log = emo.NewZone("ntf")

// Notifier interface for sending messages.
type Notifier interface {
	Notify(message string) error
}

// New selects the Notifier type depending on the endpoint pattern.
func New(endpoint string) Notifier {
	switch endpoint {
	case "":
		log.Info("empty URL => use the logger Notifier")
		return logger.NewNotifier()
	default:
		return mattermost.NewNotifier(endpoint)
	}
}
