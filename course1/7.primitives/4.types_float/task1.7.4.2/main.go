package main

import (
	"fmt"
	"time"

	"github.com/mattevans/dinero"
)

func main() {
	fmt.Println(currencyPairRate("USD", "EUR", 100.0))
}

func currencyPairRate(a, b string, amount float64) float64 {
	client := dinero.NewClient("27310dcf509547c5b5bb8932eee2ad6b", a, 20*time.Minute)
	rates, err := client.Rates.Get(b)
	if err != nil {
		fmt.Println("error with trying to get rates:", err)
		return 0.0
	}
	return amount * (*rates)
}
