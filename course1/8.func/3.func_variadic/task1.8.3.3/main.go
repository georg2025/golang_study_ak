package main

import "fmt"

func main() {
	PrintNumbers(1, 3, 5, 1, 2, 4)
}

func PrintNumbers(numbers ...int) {
	for _, i := range numbers {
		fmt.Println(i)
	}
}
