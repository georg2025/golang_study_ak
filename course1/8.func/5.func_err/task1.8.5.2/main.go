package main

import (
	"fmt"
)

func main() {
	totalPrice, err := CheckDiscount(112.4, 45.5)
	if err != nil {
		panic(err)
	}
	fmt.Println(totalPrice)
}

func CheckDiscount(price, discount float64) (float64, error) {
	if discount > 50.0 {
		return price, fmt.Errorf("discount cant be more than 50%%")
	}
	if price <= 0.0 {
		return 0.0, fmt.Errorf("price cant be zero or less")
	}
	return price - (price * discount / 100), nil
}
