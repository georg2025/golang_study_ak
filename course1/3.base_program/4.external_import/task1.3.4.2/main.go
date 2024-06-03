package main

import (
	"fmt"

	"github.com/ksrof/gocolors"
)

func main() {
	fmt.Println(ColorizeRed("Hello World!"))
	fmt.Println(ColorizeGreen("Hello World!"))
	fmt.Println(ColorizeBlue("Hello World!"))
	fmt.Println(ColorizeYellow("Hello World!"))
	fmt.Println(ColorizeMagenta("Hello World!"))
	fmt.Println(ColorizeCyan("Hello World!"))
	fmt.Println(ColorizeWhite("Hello World!"))

}

func ColorizeRed(a string) string {
	return gocolors.Red(a, "")
}

func ColorizeGreen(a string) string {

	return gocolors.Green(a, "")
}

func ColorizeBlue(a string) string {
	return gocolors.Blue(a, "")
}

func ColorizeYellow(a string) string {

	return gocolors.Yellow(a, "")
}
func ColorizeMagenta(a string) string {
	return gocolors.Magenta(a, "")
}

func ColorizeCyan(a string) string {

	return gocolors.Cyan(a, "")
}
func ColorizeWhite(a string) string {
	return gocolors.White(a, "")
}
