package main

import (
	"net"
	"testing"
	"time"
)

func TestBroadcaster(t *testing.T) {
	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)
	client1 := client{ch: ch1}
	client2 := client{ch: ch2}

	go broadcaster()

	entering <- client1
	entering <- client2

	time.Sleep(time.Millisecond)
	messages <- "test message"

	msg := <-ch1
	if msg != "test message" {
		t.Errorf("expected 'test message', got %s", msg)
	}
	msg = <-ch2
	if msg != "test message" {
		t.Errorf("expected 'test message', got %s", msg)
	}
}

type mockConn struct {
	net.Conn
	readData  string
	writeData string
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	copy(b, m.readData)
	return len(m.readData), nil
}

func (m *mockConn) Write(b []byte) (n int, err error) {
	m.writeData = string(b)
	return len(b), nil
}

func (m *mockConn) Close() error {
	return nil
}

func (m *mockConn) RemoteAddr() net.Addr {
	return &net.IPAddr{IP: net.ParseIP("127.0.0.1")}
}

func TestHandleConn(t *testing.T) {
	conn := &mockConn{readData: "Hello"}
	go handleConn(conn)

	time.Sleep(1 * time.Millisecond)

	if conn.writeData == "" {
		t.Errorf("Expected data to be written to connection, but got empty string")
	}
}
