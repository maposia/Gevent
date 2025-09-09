package main

import (
	"github.com/maposia/gevent/calendar"
	"github.com/maposia/gevent/cmd"
	"github.com/maposia/gevent/storage"
)

func main() {
	sz := storage.NewZipStorage("calendar.zip")
	c := calendar.NewCalendar(sz)

	cli := cmd.NewCmd(c)
	cli.Run()

}
