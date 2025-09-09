package events

import (
	"errors"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/maposia/gevent/reminder"
)

type Event struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	StartAt  time.Time          `json:"startAt"`
	Priority Priority           `json:"priority"`
	Reminder *reminder.Reminder `json:"reminder"`
}

var ErrTooLateReminder = errors.New("reminder time is later than event")
var ErrWrongDate = errors.New("wrong date format")

func NewEvent(title string, dateStr string, priority Priority) (*Event, error) {
	isValidTitle := ValidateTitle(title)
	if !isValidTitle {
		return &Event{}, errors.New("new event: invalid event name format")
	}

	if err := priority.Validate(); err != nil {
		return &Event{}, err
	}

	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return &Event{}, errors.New("new event: invalid date format")
	}

	return &Event{ID: getUniqueID(), Title: title, StartAt: t, Priority: priority, Reminder: nil}, nil
}

func (e *Event) UpdateEvent(title string, date string, priority Priority) error {
	isValidTitle := ValidateTitle(title)
	if !isValidTitle {
		return errors.New("update event: title not valid")
	}

	t, err := dateparse.ParseAny(date)
	if err != nil {
		return err
	}

	if err := priority.Validate(); err != nil {
		return err
	}

	e.Title = title
	e.StartAt = t
	e.Priority = priority
	return nil
}

func (e *Event) AddReminder(message string, at string, notifier func(string, string)) error {

	t, err := dateparse.ParseAny(at)
	if err != nil {
		return fmt.Errorf("can't add reminder: %w", ErrWrongDate)
	}
	if e.StartAt.Before(t) {
		return fmt.Errorf("can't add reminder: %w", ErrTooLateReminder)
	}

	newReminder, err := reminder.NewReminder(message, t, notifier)

	if err != nil {
		return err
	}

	e.Reminder = newReminder
	e.Reminder.Start()
	return nil
}

func (e *Event) RemoveReminder() {
	e.Reminder = nil
}
