package main

import (
	"fmt"
	"sync"
)

func main() {
	chan1 := generateChan(5)
	chan2 := generateChan(5)
	mergedChan := mergeChan2(chan1, chan2)

	for i := range mergedChan {
		fmt.Println(i)
	}
}

func mergeChan(mergeTo chan int, from ...chan int) {
	go func() {
		var wg sync.WaitGroup

		for i := 0; i < len(from); i++ {
			wg.Add(1)
			ch := from[i]

			go func() {
				defer wg.Done()
				for j := range ch {
					mergeTo <- j
				}
			}()
		}
	}()
}

func mergeChan2(chans ...chan int) chan int {
	mergeTo := make(chan int)

	go func() {
		defer close(mergeTo)
		var wg sync.WaitGroup

		for i := 0; i < len(chans); i++ {
			wg.Add(1)
			ch := chans[i]

			go func() {
				defer wg.Done()

				for j := range ch {
					mergeTo <- j
				}
			}()
		}
		wg.Wait()
	}()

	return mergeTo
}

func generateChan(n int) chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}

		close(ch)
	}()

	return ch
}
