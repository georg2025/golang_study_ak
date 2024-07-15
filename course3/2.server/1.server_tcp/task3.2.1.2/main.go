package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	request := make([]byte, 1024)
	_, err := conn.Read(request)

	if err != nil {
		fmt.Println("reading request error: ", err)
		return
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		request := scanner.Text()
		parts := strings.Split(string(request), " ")

		if parts[0] == "GET" && parts[1] == "/" {
			var sb strings.Builder
			sb.WriteString("HTTP/1.1 200 OK\r\n")
			sb.WriteString("Content-Type: text/html\r\n")
			sb.WriteString("\r\n")

			html := `
			<html>
			<head>
			<title>Webserver</title>
			</head>
			<body>
			<p>hello world</p>
			</body>
			</html>
		`
			sb.WriteString(html)

			_, err = conn.Write([]byte(sb.String()))

			if err != nil {
				fmt.Println("Error with message sending: ", err)
			}
		} else {
			var sb strings.Builder
			sb.WriteString("HTTP/1.1 404 Not Found\r\n")
			sb.WriteString("\r\n")

			_, err = conn.Write([]byte(sb.String()))

			if err != nil {
				fmt.Println("Error with message sending: ", err)
			}
		}
	}

	fmt.Println("Client disconnected: ", conn.RemoteAddr().String())
}

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}
