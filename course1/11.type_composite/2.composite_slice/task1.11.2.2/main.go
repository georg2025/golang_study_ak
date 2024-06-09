package main

func main() {
}

func MaxDifference(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	min := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return max - min
}
