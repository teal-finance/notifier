// Package logger implements a Notifier interface that logs the received notifications.
package logger

import (
	"strings"
	"unicode/utf8"

	"github.com/teal-finance/emo"
)

var log = emo.NewZone("ntf")

// Notifier is an empty struct.
type Notifier struct{}

// NewNotifier creates a new logger Notifier.
func NewNotifier() Notifier {
	return Notifier{}
}

// Notify prints the messages to the logs.
func (n Notifier) Notify(msg string) error {
	log.Info("LoggerNotifier:", sanitize(msg))
	return nil
}

// The code points in the surrogate range are not valid for UTF-8.
const (
	surrogateMin = 0xD800
	surrogateMax = 0xDFFF
)

// sanitize replaces control codes by the tofu symbol
// and invalid UTF-8 codes by the replacement character.
// sanitize can be used to prevent log injection.
//
// Inspired from:
// - https://wikiless.org/wiki/Replacement_character#Replacement_character
// - https://graphicdesign.stackexchange.com/q/108297
func sanitize(str string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case r < 32, r == 127: // The .notdef character is often represented by the empty box (tofu)
			return '􏿮' // to indicate a valid but not rendered character.
		case surrogateMin <= r && r <= surrogateMax, utf8.MaxRune < r:
			return '�' // The replacement character U+FFFD indicates an invalid UTF-8 character.
		}
		return r
	}, str)
}
