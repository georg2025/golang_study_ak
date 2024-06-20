package main

func main() {
}

func sum(numbers [8]int) int {
	answer := 0
	for _, i := range numbers {
		answer += i
	}
	return answer
}

func average(numbers [8]int) float64 {
	summ := sum(numbers)
	return float64(summ) / 8.0
}

func averageFloat(numbers [8]float64) float64 {
	summ := 0.0
	for _, i := range numbers {
		summ += i
	}
	return summ / 8.0
}

func reverse(numbers [8]int) [8]int {
	x := 7
	for i := 0; i < x; i++ {
		numbers[i], numbers[x] = numbers[x], numbers[i]
		x -= 1
	}
	return numbers
}
