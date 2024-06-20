package main

func main() {
}

func InsertToStart(xs []int, x ...int) []int {
	if len(x) == 0 {
		return xs
	}

	newSlice := make([]int, 0, len(xs)+len(x))
	newSlice = append(newSlice, x...)
	newSlice = append(newSlice, xs...)
	return newSlice
}
