package main

import (
	"fmt"
	mymath "test/test"
)

func main() {
	x := 4.4
	y := 12.3
	fmt.Println(mymath.Sqrt(x))
	fmt.Println(mymath.Ceil(x))
	fmt.Println(mymath.Floor(x))
	fmt.Println(mymath.Pow(x, y))
	fmt.Println(mymath.Max(x, y))
	fmt.Println(mymath.Min(x, y))
}
