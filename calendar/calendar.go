package calendar

import (
	"encoding/json"
	"errors"

	"github.com/maposia/go_calendar/events"
	"github.com/maposia/go_calendar/logger"
	"github.com/maposia/go_calendar/storage"
)

type NotificationType string

const (
	Info   NotificationType = "info"
	Error  NotificationType = "error"
	System NotificationType = "system"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        storage.Store
	Notification   chan string
}

var ErrEventNotFound = errors.New("event not found")

func NewCalendar(s storage.Store) *Calendar {

	return &Calendar{calendarEvents: make(map[string]*events.Event), storage: s, Notification: make(chan string)}
}

func (c *Calendar) AddEvent(title string, date string, priority events.Priority) (*events.Event, error) {
	event, err := events.NewEvent(title, date, priority)
	if err != nil {
		return &events.Event{}, err
	}

	c.calendarEvents[event.ID] = event

	c.Notify("event "+event.Title+" added successfully", "Info")
	return event, nil
}

func (c *Calendar) DeleteEvent(key string) error {
	event, ok := c.calendarEvents[key]
	if !ok {
		return errors.New("delete event: deletion not possible, event not found")
	}

	c.Notify("event "+event.Title+" successfully removed", "Info")
	delete(c.calendarEvents, key)
	return nil
}

func (c *Calendar) EditEvent(key string, title string, date string, priority events.Priority) error {
	e, exist := c.calendarEvents[key]
	if !exist {
		return errors.New("event edit: event not found")
	}

	err := e.UpdateEvent(title, date, priority)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) GetEvents() map[string]*events.Event {
	if len(c.calendarEvents) == 0 {
		return map[string]*events.Event{}
	}
	return c.calendarEvents
}

func (c *Calendar) SetEventReminder(key string, message string, at string) {
	e, exist := c.calendarEvents[key]
	if !exist {
		c.Notify("can't set reminder: "+ErrEventNotFound.Error(), "Error")
	} else {
		if err := e.AddReminder(message, at, c.Notify); err != nil {
			c.Notify(err.Error(), "Error")
		}
	}

}

func (c *Calendar) RemoveEventReminder(key string) error {
	e, ok := c.calendarEvents[key]
	if !ok {
		return errors.New("remove reminder: event not found")
	}
	e.RemoveReminder()
	c.Notify("notification successfully removed", "Info")

	return nil
}

func (c *Calendar) Notify(msg string, notificationType string) {

	switch notificationType {
	case "Error":
		logger.Error(msg)
		break
	case "Info":
		logger.Info(msg)
		break
	case "System":
		logger.System(msg)
		break
	default:
		logger.Info(msg)
		break
	}

	c.Notification <- msg
}

func (c *Calendar) Save() error {
	data, err := json.Marshal(c.calendarEvents)
	if err != nil {
		return err
	}
	err = c.storage.Save(data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c.calendarEvents)
	if err != nil {
		return err
	}

	return nil
}
