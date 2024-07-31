package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

type client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

var (
	// chanel for incomming clients
	entering = make(chan client)
	// chanel for clients are out
	leaving = make(chan client)
	// chanel for messenges
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

// broadcaster send incomming messages to all clients
// and looking for incomming and outcomming clients
func broadcaster() {
	// here we got all current clients
	clients := make(map[client]bool)
	var mutex sync.Mutex
	go clientsAdding(clients, &mutex)
	go clientDeleting(clients, &mutex)

	for message := range messages {
		mutex.Lock()

		for client := range clients {
			client.ch <- message
		}

		mutex.Unlock()
	}
}

func clientsAdding(clientMap map[client]bool, mutex *sync.Mutex) {
	for clients := range entering {
		mutex.Lock()
		clientMap[clients] = true
		mutex.Unlock()
	}
}

func clientDeleting(clientMap map[client]bool, mutex *sync.Mutex) {
	for clients := range leaving {
		mutex.Lock()
		delete(clientMap, clients)
		close(clients.ch)
		mutex.Unlock()
	}
}

// handleConn process incomming messages from client
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	cli := client{conn, who, ch}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli
	input := bufio.NewScanner(conn)

	for input.Scan() {
		messages <- (who + ": " + input.Text())
	}

	messages <- who + " has left"
	conn.Close()
}

// clientWriter send messages to current client
func clientWriter(conn net.Conn, ch <-chan string) {
	for message := range ch {
		_, err := conn.Write([]byte(message))

		if err != nil {
			return
		}
	}
}
