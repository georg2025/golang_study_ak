package main

import "fmt"

func main() {
	number := 11
	changeInt(&number)
	floatNumber := 1.3
	changeFloat(&floatNumber)
	ourString := "Hello, world!"
	changeString(&ourString)
	ourBool := true
	changeBool(&ourBool)
	fmt.Println(number)
	fmt.Println(floatNumber)
	fmt.Println(ourString)
	fmt.Println(ourBool)
}

func changeInt(a *int) {
	*a = 20
}

func changeFloat(b *float64) {
	*b = 6.28
}

func changeString(c *string) {
	*c = "Goodbye, world!"
}

func changeBool(d *bool) {
	*d = false
}
