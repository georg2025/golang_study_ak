package main

import (
	"io"
	"log"
	"os"
)

type RemoteLogger struct {
	Address string
}

func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}

type Logger interface {
	Log(string) error
}

type ConsoleLogger struct {
	Writer io.Writer
}

type FileLogger struct {
	File string
}

func (logger *ConsoleLogger) Log(info string) error {
	_, err := logger.Writer.Write([]byte(info))
	return err
}

func (logger *FileLogger) Log(info string) error {
	file, err := os.OpenFile(logger.File, os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(info))
	return err
}

func main() {
	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
	fileLogger := &FileLogger{File: "logger.txt"}

	loggers := []Logger{consoleLogger, fileLogger}
	LogAll(loggers, "This is a test log message.")
}
