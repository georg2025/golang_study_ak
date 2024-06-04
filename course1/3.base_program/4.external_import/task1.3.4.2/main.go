// This library was supposed to be written full with gocolors lib. Unfortunately, it is unable
// to make RGB color with this lib, so, it was made manually with the help of fmt.Sprintf

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
	fmt.Println(ColorizeCustom("Hello World!", 100, 255, 50))

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

func ColorizeCustom(a string, r, g, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, a)
}
