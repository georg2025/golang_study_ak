package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Floor(6.8))
}

func Floor(x float64) float64 {
	return math.Floor(x)
}
