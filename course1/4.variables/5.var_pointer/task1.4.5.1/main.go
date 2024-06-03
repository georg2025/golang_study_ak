package main

import "fmt"

func main() {
	number := 11
	changeInt(&number)
	floatnumber := 1.3
	changeFloat(&floatnumber)
	ourstring := "Hello, world!"
	changeString(&ourstring)
	ourbool := true
	changeBool(&ourbool)
	fmt.Println(number)
	fmt.Println(floatnumber)
	fmt.Println(ourstring)
	fmt.Println(ourbool)

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
