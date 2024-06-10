package main

func main() {
}

func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	if len(xs) <= idx || idx < 0 {
		return []int{}
	}

	xs = append(xs[:idx+1], x...)
	return xs
}
