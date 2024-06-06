package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypotenuse(4.3, 7.5))
}

func hypotenuse(a, b float64) float64 {
	return math.Hypot(a, b)
}
