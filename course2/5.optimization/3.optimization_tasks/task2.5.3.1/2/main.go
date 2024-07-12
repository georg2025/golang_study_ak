// https://leetcode.com/problems/palindrome-number/description/
package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	numbers := []int{}
	for x > 0 {
		z := x % 10
		numbers = append(numbers, z)
		x = x / 10
	}

	startFlag := 0
	endFlag := len(numbers) - 1

	for startFlag < endFlag {
		if numbers[startFlag] != numbers[endFlag] {
			return false
		}

		startFlag++
		endFlag--
	}

	return true
}

func main() {
}
