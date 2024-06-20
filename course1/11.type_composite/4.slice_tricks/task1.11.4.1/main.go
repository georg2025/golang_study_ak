package main

func main() {
}

func Cut(xs []int, start, end int) []int {
	if len(xs) == 0 || start > end || start < 0 || end >= len(xs) {
		return []int{}
	}
	newSlice := xs[start : end+1]
	return newSlice
}
