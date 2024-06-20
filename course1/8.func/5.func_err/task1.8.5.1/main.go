package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := CalculatePercentageChange("65.4", "77.8")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	initialFloat, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0.0, err
	}
	finalFloat, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0.0, err
	}
	if initialFloat == 0.0 {
		fmt.Println("Cant count, cause initial Value is zero")
		return 0.0, nil
	}
	return (finalFloat/initialFloat - 1) * 100, nil
}
