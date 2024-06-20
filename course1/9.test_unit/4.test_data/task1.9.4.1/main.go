package main

func main() {
}

func average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0.0
	}
	totalCount := 0.0
	for _, i := range xs {
		totalCount += i
	}
	return totalCount / float64(len(xs))
}
