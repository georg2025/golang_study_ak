package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindMaxAndMin(4, 2, 5, 76, 8, 3, 8, 0, 7, 3, 2, 1, -3))
}

func FindMaxAndMin(n ...int) (int, int) {
	if len(n) == 0 {
		return 0, 0
	}
	max := n[0]
	min := n[0]

	for _, i := range n {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}

	return max, min
}
