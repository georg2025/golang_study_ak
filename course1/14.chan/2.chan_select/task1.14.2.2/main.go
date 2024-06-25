package main

import (
	"time"
)

func main() {
}

func timeout(timeout time.Duration) func() bool {
	return func() bool {
		done := make(chan bool)
		defer close(done)

		go func() {
			time.Sleep(5 * time.Second)
			done <- true
		}()

		time.Sleep(timeout)

		select {
		case <-done:
			return true
		default:
			return false
		}
	}
}
