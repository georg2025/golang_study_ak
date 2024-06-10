package main

func main() {
}

func RemoveIDX(xs []int, idx int) []int {
	if len(xs) == 0 {
		return []int{}
	}
	if idx < 0 || idx >= len(xs) {
		return xs
	}
	if idx == len(xs)-1 {
		xs = xs[:idx]
		return xs
	}
	xs = append(xs[:idx], xs[idx+1:]...)
	return xs
}
