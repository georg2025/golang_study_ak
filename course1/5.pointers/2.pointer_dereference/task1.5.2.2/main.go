package main

import (
	"fmt"
	"strings"
)

func main() {
	x := 6
	fmt.Println(Factorial(&x))
	y := "abrsrba"
	fmt.Println(isPalindrome(&y))
	slice := []int{3, 5, 1, 6, 7, 8, 3, 4, 3, 4, 5, 67, 3, 3}
	number := 3
	fmt.Println(CountOccurrences(&slice, &number))
	z := "Hello World!"
	fmt.Println(ReverseString(&z))
}

func Factorial(n *int) int {
	answer := 1
	for i := *n; i > 1; i-- {
		answer *= i
	}
	return answer
}

func isPalindrome(str *string) bool {
	letters := strings.Split(*str, "")
	j := len(letters) - 1
	for i := 0; i < j; i++ {
		if letters[i] != letters[j] {
			return false
		}
		j--
	}
	return true
}

func CountOccurrences(numbers *[]int, target *int) int {
	answer := 0
	if len(*numbers) == 0 {
		return answer
	}
	for i := 0; i < len(*numbers); i++ {
		if (*numbers)[i] == *target {
			answer += 1
		}
	}
	return answer
}

func ReverseString(str *string) string {
	if len(*str) == 0 {
		return ""
	}
	letters := strings.Split(*str, "")
	builder := &strings.Builder{}
	for i := (len(letters) - 1); i >= 0; i-- {
		builder.WriteString(letters[i])
	}
	return builder.String()
}
