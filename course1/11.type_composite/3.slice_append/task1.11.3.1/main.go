package main

func main() {
}

func appendInt(xs []int, x ...int) []int {
	return append(xs, x...)
}
