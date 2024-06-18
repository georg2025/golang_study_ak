package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	fmt.Println(binaryStringToFloat("00111101101100001001001010000001"))
}

func binaryStringToFloat(binary string) float32 {
	var number uint32
	var checkingBit uint32 = 1
	numbersSlice := strings.Split(binary, "")

	for i := len(numbersSlice) - 1; i >= 0; i-- {
		if numbersSlice[i] == "1" {
			number = number | checkingBit
		}
		checkingBit = checkingBit << 1
	}

	return *(*float32)(unsafe.Pointer(&number))
}
