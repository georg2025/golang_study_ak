package main

import "fmt"

func main() {
	addThree := adder(3)
	fmt.Println(addThree(6))
	fmt.Println(addThree(533))
	fmt.Println(addThree(-2))
}

func adder(initial int) func(int) int {
	return func(x int) int { return x + initial }
}
