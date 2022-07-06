// Package logger implements a Notifier interface that logs the received notifications.
package logger

import (
	"log"
)

// Notifier is an empty struct.
type Notifier struct{}

// NewNotifier creates a new logger Notifier.
func NewNotifier() Notifier {
	return Notifier{}
}

// Notify prints the messages to the logs.
func (n Notifier) Notify(msg string) error {
	log.Print("LoggerNotifier: ", msg)
	return nil
}
