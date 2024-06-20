package main

func main() {
}

func FilterDividers(xs []int, divider int) []int {
	if len(xs) == 0 {
		return []int{}
	}

	newSlice := []int{}
	for _, i := range xs {
		if i%divider == 0 {
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}
