package main

import (
	"sort"
)

func main() {
}

func sortDescInt(array [8]int) [8]int {
	sort.Slice(array[:], func(i, j int) bool {
		return array[i] > array[j]
	})
	return array
}

func sortAscInt(array [8]int) [8]int {
	sort.Ints(array[:])
	return array
}

func sortDescFloat(array [8]float64) [8]float64 {
	sort.Float64s(array[:])
	x := 7
	for i := 0; i < x; i++ {
		array[x], array[i] = array[i], array[x]
		x -= 1
	}
	return array
}

func sortAscFloat(array [8]float64) [8]float64 {
	sort.Float64s(array[:])
	return array
}
