package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	data := NotifyEvery(ticker, 5*time.Second, "Timer alert")

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Program is over")
}

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	newChan := make(chan string)

	go func() {
		for _ = range ticker.C {
			newChan <- message
		}
	}()

	go func() {
		time.Sleep(d)
		ticker.Stop()
		close(newChan)
	}()

	return newChan
}
