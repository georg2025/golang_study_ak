package main

import "fmt"

func main() {
}

func getVariableType(variable interface{}) string {
	if variable == nil {
		return "Пустой интерфейс"
	}
	answer := ""
	switch variable.(type) {
	case int:
		answer = "int"
	case string:
		answer = "string"
	case []int:
		answer = "[]int"
	default:
		return "Cant figure out variable type"
	}
	return fmt.Sprintf("Variable type is: %s", answer)
}
