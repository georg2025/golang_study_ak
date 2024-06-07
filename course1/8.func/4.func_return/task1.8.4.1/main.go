package main

import (
	"fmt"
)

func main() {
	fmt.Println(DivideAndReemainder(33, 5))
}

func DivideAndReemainder(a, b int) (int, int) {
	if b == 0 {
		fmt.Println("Cant divide on zero")
		return 0, 0
	}
	return a / b, a % b
}
