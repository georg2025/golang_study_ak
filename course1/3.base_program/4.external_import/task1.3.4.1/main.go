package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	number1 := "5.8"
	number2 := "3.4"
	number3 := "6.49563888456"

	summ, err := DecimalSum(number1, number2)
	if err != nil {
		fmt.Println("Error = ", err)
	} else {
		fmt.Println("Summ = ", summ)
	}

	subtract, err := DecimalSubtract(number1, number2)
	if err != nil {
		fmt.Println("Error = ", err)
	} else {
		fmt.Println("Subtract = ", subtract)
	}

	multiply, err := DecimalMultiply(number1, number2)
	if err != nil {
		fmt.Println("Error = ", err)
	} else {
		fmt.Println("Multiply = ", multiply)
	}

	divide, err := DecimalDivide(number1, number2)
	if err != nil {
		fmt.Println("Error = ", err)
	} else {
		fmt.Println("Divide = ", divide)
	}

	round, err := DecimalRound(number3, 7)
	if err != nil {
		fmt.Println("Error = ", err)
	} else {
		fmt.Println("Round = ", round)
	}

	isBigger, err := DecimalGreaterThan(number1, number2)
	if err != nil {
		fmt.Println("Error = ", err)
	} else {
		fmt.Println("First is bigger? It is", isBigger)
	}

	isSmaller, err := DecimalLessThan(number1, number2)
	if err != nil {
		fmt.Println("Error = ", err)
	} else {
		fmt.Println("First is smaller? It is", isSmaller)
	}

	equal, err := DecimalEqual(number1, number2)
	if err != nil {
		fmt.Println("Error = ", err)
	} else {
		fmt.Println("Numbers are equal? It is", equal)
	}
}

func DecimalSum(a, b string) (string, error) {
	first, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	second, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	answer := decimal.Sum(first, second)
	return answer.String(), nil
}

func DecimalSubtract(a, b string) (string, error) {
	first, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	second, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	answer := first.Sub(second)
	return answer.String(), nil
}

func DecimalMultiply(a, b string) (string, error) {
	first, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	second, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	answer := first.Mul(second)
	return answer.String(), nil
}

func DecimalDivide(a, b string) (string, error) {
	first, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	second, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	answer := first.Div(second)
	return answer.String(), nil
}

func DecimalRound(a string, precision int32) (string, error) {
	first, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	answer := first.Round(precision)
	return answer.String(), nil
}

func DecimalGreaterThan(a, b string) (bool, error) {
	first, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	second, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}

	answer := first.GreaterThan(second)
	return answer, nil
}

func DecimalLessThan(a, b string) (bool, error) {
	first, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	second, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}

	answer := first.LessThan(second)
	return answer, nil
}

func DecimalEqual(a, b string) (bool, error) {
	first, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	second, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}

	answer := first.Equal(second)
	return answer, nil
}
