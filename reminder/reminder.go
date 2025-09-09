package reminder

import (
	"fmt"
	"time"
)

type Reminder struct {
	Message  string               `json:"message"`
	At       time.Time            `json:"at"`
	Sent     bool                 `json:"sent"`
	Timer    *time.Timer          `json:"timer"`
	Notifier func(string, string) `json:"notifier"`
}

func NewReminder(message string, at time.Time, notifier func(string, string)) (*Reminder, error) {
	err := ValidateMessage(message)
	if err != nil {
		return &Reminder{}, fmt.Errorf("can't create new reminder: %w", err)
	}

	return &Reminder{message, at, false, nil, notifier}, nil
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}
	r.Notifier(r.Message, "Info")
	r.Sent = true
}

func (r *Reminder) Stop() {
	r.Timer.Stop()
}

func (r *Reminder) Start() {
	now := time.Now()
	targetTime := r.At.In(now.Location())

	r.Timer = time.AfterFunc(time.Until(targetTime), r.Send)

}
