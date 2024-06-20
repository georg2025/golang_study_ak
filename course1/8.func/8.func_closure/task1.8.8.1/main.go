package main

import (
	"fmt"
)

func main() {
	counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func createCounter() func() int {
	x := 0
	return func() int {
		x += 1
		return x
	}
}
