package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {

	number1 := "5.8"
	number2 := "3.4"
	number3 := "6.49563888456"

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

	multiply, error := DecimalMultiply(number1, number2)
	if error != nil {
		fmt.Println("Error = ", error)
	} else {
		fmt.Println("Multiply = ", multiply)
	}

	divide, error := DecimalDivide(number1, number2)
	if error != nil {
		fmt.Println("Error = ", error)
	} else {
		fmt.Println("Divide = ", divide)
	}

	round, error := DecimalRound(number3, 7)
	if error != nil {
		fmt.Println("Error = ", error)
	} else {
		fmt.Println("Round = ", round)
	}

	isbigger, error := DecimalGreaterThan(number1, number2)
	if error != nil {
		fmt.Println("Error = ", error)
	} else {
		fmt.Println("First is bigger? It is", isbigger)
	}

	issmaller, error := DecimalLessThan(number1, number2)
	if error != nil {
		fmt.Println("Error = ", error)
	} else {
		fmt.Println("First is smaller? It is", issmaller)
	}

	equal, error := DecimalEqual(number1, number2)
	if error != nil {
		fmt.Println("Error = ", error)
	} else {
		fmt.Println("Numbers are equal? It is", equal)
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

func DecimalMultiply(a, b string) (string, error) {

	first, error := decimal.NewFromString(a)
	if error != nil {
		return "", error
	}

	second, error := decimal.NewFromString(b)
	if error != nil {
		return "", error
	}

	answer := first.Mul(second)
	return answer.String(), nil

}

func DecimalDivide(a, b string) (string, error) {

	first, error := decimal.NewFromString(a)
	if error != nil {
		return "", error
	}

	second, error := decimal.NewFromString(b)
	if error != nil {
		return "", error
	}

	answer := first.Div(second)
	return answer.String(), nil

}

func DecimalRound(a string, precision int32) (string, error) {

	first, error := decimal.NewFromString(a)
	if error != nil {
		return "", error
	}

	answer := first.Round(precision)
	return answer.String(), nil

}

func DecimalGreaterThan(a, b string) (bool, error) {

	first, error := decimal.NewFromString(a)
	if error != nil {
		return false, error
	}

	second, error := decimal.NewFromString(b)
	if error != nil {
		return false, error
	}

	answer := first.GreaterThan(second)
	return answer, nil

}

func DecimalLessThan(a, b string) (bool, error) {

	first, error := decimal.NewFromString(a)
	if error != nil {
		return false, error
	}

	second, error := decimal.NewFromString(b)
	if error != nil {
		return false, error
	}

	answer := first.LessThan(second)
	return answer, nil

}

func DecimalEqual(a, b string) (bool, error) {

	first, error := decimal.NewFromString(a)
	if error != nil {
		return false, error
	}

	second, error := decimal.NewFromString(b)
	if error != nil {
		return false, error
	}

	answer := first.Equal(second)
	return answer, nil

}
