package main

import (
	"github.com/maposia/go_calendar/calendar"
	"github.com/maposia/go_calendar/cmd"
	"github.com/maposia/go_calendar/storage"
)

func main() {
	sz := storage.NewZipStorage("calendar.zip")
	c := calendar.NewCalendar(sz)

	cli := cmd.NewCmd(c)
	cli.Run()

}
