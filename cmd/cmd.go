package cmd

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
	"github.com/maposia/go_calendar/calendar"
	"github.com/maposia/go_calendar/logger"
	"github.com/maposia/go_calendar/storage"
)

var mu sync.Mutex

type LogEntry struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

type Cmd struct {
	calendar *calendar.Calendar
	storage  storage.Store
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return &Cmd{calendar: c, storage: nil}
}

func (c *Cmd) executor(input string) {
	c.calendar.Load()
	parts, err := shlex.Split(input)
	if err != nil {
		fmt.Println("Processing error input:", err)
		return
	}

	logger.Info("command used:" + input)

	if len(parts) == 0 {
		return
	}

	cmd := strings.ToLower(parts[0])
	args := parts[1:]

	switch cmd {
	case "add":
		c.handleAdd(args)
	case "edit":
		c.handleUpdate(args)
	case "remove":
		c.handleRemove(args)
	case "list":
		c.handleList()
	case "setremind":
		c.handleSetRemind(args)
	case "removeremind":
		c.handleRemoveReminder(args)
	case "help":
		c.handleHelp()
	case "exit":
		err := c.calendar.Save()
		if err != nil {
			fmt.Println(err)
		}
		logger.System("App stopped")
		logger.Close()
		close(c.calendar.Notification)
		os.Exit(0)
	default:
		fmt.Println("Unknown command:")
		fmt.Println("Enter 'help' for list of commands")
	}
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	if strings.Contains(d.TextBeforeCursor(), " ") {
		return []prompt.Suggest{}
	}
	suggestions := []prompt.Suggest{
		{Text: "add", Description: "Add new event"},
		{Text: "edit", Description: "Edit event"},
		{Text: "remove", Description: "Remove event"},
		{Text: "list", Description: "Show all events"},
		{Text: "setRemind", Description: "Add reminder"},
		{Text: "removeRemind", Description: "Remove reminder"},
		{Text: "help", Description: "Show help"},
		{Text: "exit", Description: "Close app"},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) Run() {
	logger.Init()

	go func() {
		for msg := range c.calendar.Notification {
			fmt.Printf("%s \n", msg)
		}
	}()

	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)

	logger.System("App started")
	p.Run()

}
