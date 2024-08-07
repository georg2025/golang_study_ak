package main

import (
	"fmt"
	"runtime"
	"time"
)

func factorialRecursive(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	factorial := 1
	for i := 2; i <= n; i++ {
		factorial *= i
	}

	return factorial
}

// выдает true, если реализация быстрее и false, если медленнее
func compareWhichFactorialIsFaster(n int) string {
	start := time.Now()
	_ = factorialRecursive(n)
	elapsedRecursive := time.Since(start)

	start = time.Now()
	_ = factorialIterative(n)
	elapsedIterative := time.Since(start)

	if elapsedIterative > elapsedRecursive {
		return "Recursive func is faster"
	}

	return "Iterative func is faster"
}

func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("Go OS/Arch:", runtime.GOOS, "/", runtime.GOARCH)
	fmt.Println("Which factorial is faster?")
	fmt.Println(compareWhichFactorialIsFaster(15))
}
