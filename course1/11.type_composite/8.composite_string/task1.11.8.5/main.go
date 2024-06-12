package main

import "strings"

func main() {
}

func ReverseString(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}
	Chars := strings.Split(str, "")
	var sb strings.Builder
	for i := len(Chars) - 1; i >= 0; i-- {
		sb.WriteString(Chars[i])
	}
	return sb.String()
}
