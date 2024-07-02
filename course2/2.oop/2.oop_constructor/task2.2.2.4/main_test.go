package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLogSystemLogsMessageUsingFileLogger(t *testing.T) {
	file, err := ioutil.TempFile("", "log.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))
	logSystem.Log("Hello, world!")

	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read from temp file: %v", err)
	}

	expected := "Hello, world!"
	if string(content) != expected {
		t.Errorf("Expected %q but got %q", expected, string(content))
	}
}

func TestLogSystemLogsEmptyMessage(t *testing.T) {
	file, err := ioutil.TempFile("", "log.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))
	logSystem.Log("")

	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read from temp file: %v", err)
	}

	expected := ""
	if string(content) != expected {
		t.Errorf("Expected %q but got %q", expected, string(content))
	}
}
