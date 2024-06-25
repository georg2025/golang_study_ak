package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Горутина завершила работу")
		stop <- true
	}()

	timer := time.NewTimer(3 * time.Second)

	data := NotifyOnTimer(timer, stop)

	for v := range data {
		fmt.Println(v)
	}
}

func NotifyOnTimer(timer *time.Timer, stop chan bool) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		select {
		case <-stop:
			ch <- "Горутина завершила работу раньше, чем таймер сработал"
		case <-timer.C:
			ch <- "Горутина не успела завершить работу к моменту срабатывания таймера"
		}
	}()

	return ch
}
