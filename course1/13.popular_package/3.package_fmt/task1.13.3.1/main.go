package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
}

func generateMathString(operands []int, operator string) string {
	if len(operands) == 0 {
		return ""
	}

	var sb strings.Builder
	for range operands {
		sb.WriteString("%v")
		sb.WriteString(" ")
		sb.WriteString(operator)
		sb.WriteString(" ")
	}

	var values []interface{}
	for _, i := range operands {
		values = append(values, i)
	}

	answer := operands[0]

	if len(operands) == 1 {
		return fmt.Sprintf("%d = %d", answer, answer)
	}

	switch operator {
	case "+":
		for i := 1; i < len(operands); i++ {
			answer += operands[i]
		}
	case "*":
		for i := 1; i < len(operands); i++ {
			answer *= operands[i]
		}
	case "-":
		for i := 1; i < len(operands); i++ {
			answer -= operands[i]
		}
	case "/":
		for i := 1; i < len(operands); i++ {
			answer /= operands[i]
		}
	default:
		return ""
	}

	resultFmt := sb.String()
	resultFmt = resultFmt[:len(resultFmt)-2]
	resultFmt += "= "
	resultFmt += strconv.Itoa(answer)
	result := fmt.Sprintf(resultFmt, values...)
	return result
}
