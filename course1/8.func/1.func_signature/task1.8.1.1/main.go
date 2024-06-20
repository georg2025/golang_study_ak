package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CalculateCircleArea(5.4))
	fmt.Println(CalculateRectangleArea(5.4, 6.6))
	fmt.Println(CalculateTriangleArea(5.4, 7.5))
}

var CalculateCircleArea = func(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

var CalculateRectangleArea = func(width, height float64) float64 {
	return width * height
}

var CalculateTriangleArea = func(base, height float64) float64 {
	return base * height / 2
}
