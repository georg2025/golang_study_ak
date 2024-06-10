package main

func main() {
}

func Cut(x []int, start, end int) []int {
	if len(x) == 0 || start > end || start < 0 || end >= len(x) {
		return []int{}
	}
	return x[start : end+1]
}
