package main

import (
	"fmt"

	"github.com/mewzax/gocolors"
)

func main() {
	fmt.Println(ColorizeRed("Hello World!"))
	fmt.Println(ColorizeGreen("Hello World!"))
	fmt.Println(ColorizeBlue("Hello World!"))
	fmt.Println(ColorizeYellow("Hello World!"))
	fmt.Println(ColorizeMagenta("Hello World!"))
	fmt.Println(ColorizeCyan("Hello World!"))
	fmt.Println(ColorizeWhite("Hello World!"))
	fmt.Println(ColorizeCustom("Hello World!", 255, 50, 50))
}

func ColorizeRed(a string) string {
	return gocolors.Colorize(gocolors.Red, a)
}

func ColorizeGreen(a string) string {
	return gocolors.Colorize(gocolors.Green, a)
}

func ColorizeBlue(a string) string {
	return gocolors.Colorize(gocolors.Blue, a)
}

func ColorizeYellow(a string) string {
	return gocolors.Colorize(gocolors.Yellow, a)
}

func ColorizeMagenta(a string) string {
	return gocolors.Colorize(gocolors.Magenta, a)
}

func ColorizeCyan(a string) string {
	return gocolors.Colorize(gocolors.Cyan, a)
}

func ColorizeWhite(a string) string {
	return gocolors.Colorize(gocolors.White, a)
}

func ColorizeCustom(a string, r, g, b uint8) string {
	return gocolors.Colorize(gocolors.RGB(int(r), int(g), int(b)), a)
}
