package main

func main() {
}

func bitwiseXOR(n, res int) int {
	return n ^ res
}

func findSingleNumber(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = bitwiseXOR(numbers[i], result)
	}
	return result
}
