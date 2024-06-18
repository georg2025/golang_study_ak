package main

import "fmt"

func main() {
}

func getVariableType(variable interface{}) string {
	if variable == nil {
		return "Пустой интерфейс"
	}
	return fmt.Sprintf("Variable type is: %T", variable)
}
