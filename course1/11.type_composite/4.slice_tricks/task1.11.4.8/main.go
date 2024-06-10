package main

func main() {
}

func Shift(xs []int) (int, []int) {
	// this are two special cases: if len = 0 or 1
	if len(xs) == 0 {
		return 0, []int{}
	}
	if len(xs) == 1 {
		return xs[0], xs
	}
	// If len is more than 2, we just need to make a new slice with 1 elem and cap == len of xs
	newSlice := make([]int, 1, len(xs))
	// Our first elem now became last elem from previous slice
	newSlice[0] = xs[len(xs)-1]
	// And now we append all last elements
	newSlice = append(newSlice, xs[:len(xs)-1]...)
	return newSlice[1], newSlice
}
