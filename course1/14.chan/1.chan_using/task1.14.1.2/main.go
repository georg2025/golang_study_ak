package main

import (
	"fmt"
	"math/rand"
)

func main() {
	data := generateData(10)

	for num := range data {
		fmt.Println(num)
	}
}

func generateData(n int) chan int {
	chanel := make(chan int)

	go func() {
		defer close(chanel)
		for i := 0; i < n; i++ {
			chanel <- rand.Int()
		}
	}()

	return chanel
}
