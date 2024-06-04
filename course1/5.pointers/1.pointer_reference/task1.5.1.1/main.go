package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println(*Add(5, 6))
	slice := []int{14, 4, 2, 6, 9, 2, 7}
	fmt.Println(*Max(slice))
	fmt.Println(*IsPrime(4398042316799))
	strings := []string{"hello", " ", "world", "!"}
	fmt.Println(*ConcatenateStrings(strings))
}

func Add(a, b int) *int {
	var answer *int = new(int)
	*answer = a + b
	return answer
}

func Max(numbers []int) *int {
	var answer *int = new(int)
	if len(numbers) == 0 {
		*answer = 0
		return answer
	}
	*answer = numbers[0]
	for _, i := range numbers {
		if *answer < i {
			*answer = i
		}
	}
	return answer
}

func IsPrime(number int) *bool {
	var answer *bool = new(bool)
	*answer = big.NewInt(int64(number)).ProbablyPrime(0)
	return answer
}

func ConcatenateStrings(strs []string) *string {
	var answer *string = new(string)
	builder := &strings.Builder{}
	for _, i := range strs {
		builder.WriteString(i)
	}
	*answer = builder.String()
	return answer
}
