package main

func main() {
}

func InsertToStart(xs []int, x ...int) []int {
	if len(x) == 0 {
		return xs
	}
	newSlice := []int{}
	newSlice = append(newSlice, x...)
	newSlice = append(newSlice, xs...)
	return newSlice
}
