package main

import (
	"fmt"
)

func main() {
	fmt.Println(Factorial(6))
}

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
