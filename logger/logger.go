package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger   *log.Logger
	errorLogger  *log.Logger
	systemLogger *log.Logger
	logFile      *os.File
)

func newLogger(file *os.File, prefix string) *log.Logger {
	return log.New(file, prefix, log.Ldate|log.Ltime)
}

func Init() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("error init logger: %s", err)
	}

	infoLogger = newLogger(logFile, "INFO:")
	errorLogger = newLogger(logFile, "ERROR:")
	systemLogger = newLogger(logFile, "SYSTEM:")
}

func Info(msg string) {
	infoLogger.Output(2, msg)
}

func Error(msg string) {
	errorLogger.Output(2, msg)
}

func System(msg string) {
	systemLogger.Output(2, msg)
}

func Close() {
	logFile.Close()
}
