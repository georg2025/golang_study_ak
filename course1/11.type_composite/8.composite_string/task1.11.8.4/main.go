package main

import "strings"

func main() {
}

func concatStrings(xs ...string) string {
	if len(xs) == 0 {
		return ""
	}
	if len(xs) == 1 {
		return xs[0]
	}
	var sb strings.Builder
	for _, i := range xs {
		sb.WriteString(i)
	}
	return sb.String()
}
