package main

import "fmt"

func main() {
}

func Pop(xs []int) (int, []int) {
	if len(xs) == 0 {
		fmt.Println("Empty slice - nothing to pop")
		return 0, []int{}
	}
	if len(xs) == 1 {
		fmt.Println("only one element: ", xs[0])
		return xs[0], []int{}
	}
	fmt.Printf("Значение: %d, Новый срез: %v", xs[0], xs[1:])
	return xs[0], xs[1:]
}
