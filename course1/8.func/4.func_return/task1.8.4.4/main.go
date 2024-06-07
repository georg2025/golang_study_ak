package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CalculatePercentageChange(11.4, 8.3))
}

func CalculatePercentageChange(initialValue, finalValue float64) float64 {
	if initialValue == 0.0 {
		fmt.Println("cant calculate, cause first value is zero")
		return 0.0
	}
	return (math.Round(((finalValue / initialValue) - 1.0) * 10000)) / 100
}
