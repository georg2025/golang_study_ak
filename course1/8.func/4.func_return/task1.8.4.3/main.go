package main

import "fmt"

func main() {
	fmt.Println(CalculateStockValue(4.5, 13))
}

func CalculateStockValue(price float64, quantity int) (float64, float64) {
	return price, price * float64(quantity)
}
