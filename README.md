# Go Calendar 📅

A command-line calendar application built with Go for managing events and reminders.

---

## 📋 Overview

**Go Calendar** is a CLI application that allows you to add, view, and manage events through an interactive command-line interface.

---

## 🚀 Features

- ✅ Add events to the calendar
- 📅 View events by date
- 💾 Store data in a ZIP archive
- 🔔 Reminder system
- 📝 Interactive command-line interface
---

## 🛠 Technologies

- **Go 1.24.1+**
- **Dependencies:**
  - [`github.com/araddon/dateparse`](https://github.com/araddon/dateparse) — date parsing
  - [`github.com/c-bata/go-prompt`](https://github.com/c-bata/go-prompt) — interactive CLI
  - [`github.com/google/shlex`](https://github.com/google/shlex) — command-line parsing
  - [`github.com/google/uuid`](https://github.com/google/uuid) — unique ID generation

---

## 📁 Project Structure

```bash
gevent/
├── calendar/   # Core calendar logic
├── cmd/        # Command-line interface
├── events/     # Event management
├── logger/     # Logging system
├── reminder/   # Reminder system
├── storage/    # Data storage
├── main.go     # Application entry point
└── go.mod      # Go module
```

## Clone the repository

git clone https://github.com/maposia/Gevent.git
cd gevent

## Install dependencies

go mod download

## Run the application

go run main.go