package main

import (
	"fmt"
)

func main() {
	defer recoverFromPanic()
	myPanic()
}

func myPanic() {
	panic("my panic message")
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
