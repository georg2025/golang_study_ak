package main

import "strings"

func main() {
}

func ReverseString(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}

	chars := strings.Split(str, "")
	var sb strings.Builder

	for i := len(chars) - 1; i >= 0; i-- {
		sb.WriteString(chars[i])
	}

	return sb.String()
}
