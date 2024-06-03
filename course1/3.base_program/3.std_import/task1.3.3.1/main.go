package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sqrt(25))
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
