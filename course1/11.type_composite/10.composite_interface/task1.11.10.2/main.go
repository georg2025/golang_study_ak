package main

import "strings"

func main() {
}

func Concat(xs ...interface{}) interface{} {
	var sb strings.Builder
	for _, i := range xs {
		switch i := i.(type) {
		case string:
			if len(i) > 0 {
				sb.WriteString(i)
			}
		default:
			continue
		}
	}
	return sb.String()
}

func Sum(xs ...interface{}) interface{} {
	answer := 0
	for _, i := range xs {
		switch i := i.(type) {
		case int:
			answer += i
		default:
			continue
		}
	}
	return answer
}

func Operate(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}
