package main

import (
	"fmt"
	"time"
)

type sema chan struct{}

func New(n int) sema {
	new := make(sema, n)
	return new
}

func (s sema) Inc(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}

func (s sema) Dec(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	sem := New(1)

	for _, num := range numbers {
		sem.Dec(1)
		go func(n int) {
			fmt.Println(n)
			sem.Inc(1)
		}(num)
	}

	time.Sleep(time.Millisecond)
}
