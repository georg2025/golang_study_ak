package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ConcatenateStrings("-", "I", "am", "your", "father", ",", "Luke"))
}

func ConcatenateStrings(sep string, str ...string) string {
	builder := &strings.Builder{}
	oddStrings := []string{}
	builder.WriteString("even: ")
	if len(str) == 0 {
		return "No even, no odds"
	}

	builder.WriteString(str[0])
	for i := 1; i < len(str); i++ {
		if i%2 == 0 {
			builder.WriteString(sep)
			builder.WriteString(str[i])
		} else {
			oddStrings = append(oddStrings, str[i])
		}
	}

	if len(str) == 1 {
		builder.WriteString(" no odd")
		return builder.String()
	}

	builder.WriteString(", odd: ")
	builder.WriteString(oddStrings[0])
	for i := 1; i < len(oddStrings); i++ {
		builder.WriteString(sep)
		builder.WriteString(oddStrings[i])
	}

	return builder.String()
}
