package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {

	number1 := "5.8"
	number2 := "3.5"

	summ, error := DecimalSum(number1, number2)
	if error != nil {
		fmt.Println("Error = ", error)
	} else {
		fmt.Println("Summ = ", summ)
	}

	subtract, error := DecimalSubtract(number1, number2)
	if error != nil {
		fmt.Println("Error = ", error)
	} else {
		fmt.Println("Subtract = ", subtract)
	}

}

func DecimalSum(a, b string) (string, error) {

	first, error := decimal.NewFromString(a)
	if error != nil {
		return "", error
	}

	second, error := decimal.NewFromString(b)
	if error != nil {
		return "", error
	}

	answer := decimal.Sum(first, second)
	return answer.String(), nil

}

func DecimalSubtract(a, b string) (string, error) {

	first, error := decimal.NewFromString(a)
	if error != nil {
		return "", error
	}

	second, error := decimal.NewFromString(b)
	if error != nil {
		return "", error
	}

	answer := first.Sub(second)
	return answer.String(), nil

}
