package main

func main() {
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	currentNumber := 1
	previousNumber := 0
	for i := 2; i <= n; i++ {
		previousNumber, currentNumber = currentNumber, previousNumber+currentNumber
	}
	return currentNumber
}
