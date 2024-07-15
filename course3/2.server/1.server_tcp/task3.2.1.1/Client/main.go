package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// connecting to server
	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	// start goroutine, that will be reading all messages and print them in console
	go clientReader(conn)
	// reading stdin messages and sending them to server
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		message := scanner.Text()
		_, err := conn.Write([]byte(message + "\n"))

		if err != nil {
			fmt.Println("error with writing data")
			os.Exit(1)
		}
	}
}

// clientReader print all messages from server
func clientReader(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("error with data reading: ", err)
			return
		}

		fmt.Println(string(buffer[:n]))
	}
}
