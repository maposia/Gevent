package cmd

import (
	"fmt"
	"strings"

	"github.com/maposia/go_calendar/events"
)

func (c *Cmd) handleAdd(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Format: add \"id\" \"event name\" \"date and time\" \"priority\"")
		return
	}

	fmt.Println(parts)

	title := parts[0]
	date := parts[1]
	priority := events.Priority(parts[2])

	e, err := c.calendar.AddEvent(title, date, priority)
	if err != nil {
		fmt.Println("Error adding event:", err)
		return
	} else {
		fmt.Println("Event Added", e.Title)
	}
}

func (c *Cmd) handleUpdate(parts []string) {
	if len(parts) != 4 {
		fmt.Println("Format: edit \"id\" \"event name\" \"date and time\" \"priority\"")
		return
	}
	id := parts[0]
	title := parts[1]
	date := parts[2]
	priority := events.Priority(parts[3])

	if err := c.calendar.EditEvent(id, title, date, priority); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Event", title, "has been updated")
	}

}

func (c *Cmd) handleRemove(parts []string) {
	if len(parts) != 1 {
		fmt.Println("Format: remove \"id\"")
		return
	}
	if err := c.calendar.DeleteEvent(parts[0]); err != nil {
		fmt.Printf("Error deleting:\n %s \n", err)
	}
}

func (c *Cmd) handleSetRemind(parts []string) {
	if len(parts) != 3 {
		fmt.Println("Format: setremind \"id\" \"message\" \"date and time\"")
		return
	}

	id := parts[0]
	message := parts[1]
	date := parts[2]

	c.calendar.SetEventReminder(id, message, date)
}

func (c *Cmd) handleRemoveReminder(parts []string) {
	if len(parts) != 1 {
		fmt.Println("Format: removeremind \"id\"")
		return
	}

	id := parts[0]

	if err := c.calendar.RemoveEventReminder(id); err != nil {
		fmt.Printf("Error deleting:\n %s \n", err)
	}
}

func (c *Cmd) handleList() {
	e := c.calendar.GetEvents()
	if len(e) == 0 {
		fmt.Println("The event list is empty")
	}
	fmt.Println("Show all events:")
	for _, v := range e {
		fmt.Println(v.ID, v.Title, v.StartAt.Format("02-01-2006 15:04"), v.Priority)
	}
}

func (c *Cmd) handleHelp() {
	fmt.Println("\n📅 Calendar - Command Reference")
	fmt.Println("=" + strings.Repeat("=", 40))

	fmt.Println("\n📝 Event Management:")
	fmt.Println("  add <title> <date> <priority>        - Add new event")
	fmt.Println("  list                                 - Show all events")
	fmt.Println("  update <id> <title> <date> <priority> - Update existing event")
	fmt.Println("  delete <id>                          - Delete event")

	fmt.Println("\n⏰ Reminder Management:")
	fmt.Println("  reminder add <id> <message> <time>   - Add reminder to event")
	fmt.Println("  reminder remove <id>                 - Remove reminder from event")

	fmt.Println("\n🔧 System Commands:")
	fmt.Println("  help                                 - Show this help")
	fmt.Println("  exit                                 - Exit application")

	fmt.Println("\n📖 Usage Examples:")
	fmt.Println("  add \"Meeting with client\" \"2025-12-25 14:00\" high")
	fmt.Println("  reminder add 1 \"Prepare presentation\" \"2025-12-25 13:00\"")

	fmt.Println("\n💡 Date Formats:")
	fmt.Println("  2024-12-25                           - Date only")
	fmt.Println("  2024-12-25 14:30                     - Date and time")

	fmt.Println("\n🎯 Priority Levels:")
	fmt.Println("  low                                  - Low priority")
	fmt.Println("  medium                               - Medium priority")
	fmt.Println("  high                                 - High priority")
}
