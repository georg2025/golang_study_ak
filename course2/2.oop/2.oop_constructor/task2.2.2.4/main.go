package main

import (
	"io"
)

// Logger interface
type Logger interface {
	Log(message string)
}

type FileLogger struct {
	file io.ReadWriter
}

type ConsoleLogger struct {
	out io.ReadWriter
}

func (logger ConsoleLogger) Log(message string) {
	logger.out.Write([]byte(message))
}

func (logger FileLogger) Log(message string) {
	logger.file.Write([]byte(message))
}

type LogSystem struct {
	logger Logger
}

// LogOption functional option type
type LogOption func(*LogSystem)

func WithLogger(logger Logger) LogOption {
	return func(ls *LogSystem) {
		ls.logger = logger
	}
}

func NewLogSystem(options ...LogOption) *LogSystem {
	logSystem := &LogSystem{}

	for _, option := range options {
		option(logSystem)
	}

	return logSystem
}

func (logSystem LogSystem) Log(message string) {
	logSystem.logger.Log(message)
}

func main() {
}
