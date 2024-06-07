package main

import (
	"fmt"
)

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))
	fmt.Println(MathOperate(Mul, 5, 8, 3))
	fmt.Println(MathOperate(Sub, 22, 1, 3))
}

func Sum(a ...int) int {
	result := 0
	for _, i := range a {
		result += i
	}
	return result
}

func Mul(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	result := 1
	for _, i := range a {
		result *= i
	}
	return result
}

func Sub(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	result := a[0]
	for i := 1; i < len(a); i++ {
		result -= a[i]
	}
	return result
}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}
