package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		currentTime := time.Now()
		fmt.Print("\033[h\033[2J\r\033[F")
		hour, min, sec := currentTime.Clock()
		year, month, day := currentTime.Date()
		fmt.Printf("Current time: %d:%d:%d\n", hour, min, sec)
		fmt.Printf("Current date: %d-%d-%d", year, month, day)
		time.Sleep(time.Second)
	}
}
