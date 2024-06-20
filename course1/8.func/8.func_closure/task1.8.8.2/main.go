package main

import (
	"fmt"
)

func main() {
	m := multiplier(3.6)
	fmt.Println(m(9.0))
	fmt.Println(m(6.3))
	fmt.Println(m(11.56))
}

func multiplier(factor float64) func(float64) float64 {
	return func(x float64) float64 { return x * factor }
}
