package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CompareRoundedValues(4.66656, 4.86666, 3))
}

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	helpNumber := float64(math.Pow10(decimalPlaces))
	newA := (math.Round(a * helpNumber)) / helpNumber
	newB := (math.Round(b * helpNumber)) / helpNumber
	if newA == newB {
		return true, 0.0
	}

	if newA > newB {
		return false, newA - newB
	}

	return false, newB - newA
}
