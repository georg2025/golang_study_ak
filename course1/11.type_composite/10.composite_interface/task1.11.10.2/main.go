package main

import "strings"

func main() {
}

func Concat(xs ...interface{}) interface{} {
	var sb strings.Builder
	for _, i := range xs {
		j := i.(string)
		sb.WriteString(j)
	}
	return sb.String()
}

func Sum(xs ...interface{}) interface{} {
	answer := 0
	for _, i := range xs {
		j := i.(int)
		answer += j
	}
	return answer
}

func Operate(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}
