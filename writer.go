// Package logger implements a Notifier interface that logs the received notifications.
package notifier

import (
	"io"
)

type notifierWriter struct {
	w io.Writer
}

// Notify prints the messages to the logs.
func (n notifierWriter) Notify(msg string) error {
	_, err := n.w.Write([]byte(msg))
	return err
}

// WithWriter creates simple notifier writing messages on w.
func WithWriter(w io.Writer) Notifier {
	return notifierWriter{
		w: w,
	}
}
