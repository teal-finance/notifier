package notifier

// Notifier interface for sending messages.
type Notifier interface {
	Notify(message string) error
}
