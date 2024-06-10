package main

func main() {
}

func RemoveExtraMemory(xs []int) []int {
	if len(xs) == cap(xs) {
		return xs
	}
	newSlice := make([]int, 0, len(xs))
	newSlice = append(newSlice, xs...)
	return newSlice
}
