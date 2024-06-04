package main

import (
	"fmt"
	"strings"
)

func main() {
	x := 21
	mutate(&x)
	fmt.Println(x)
	string := "Hello world!"
	ReverseString(&string)
	fmt.Println(string)
}

func mutate(a *int) {
	*a = 42
}

func ReverseString(str *string) {
	letters := strings.Split(*str, "")
	builder := &strings.Builder{}
	for i := (len(letters) - 1); i >= 0; i-- {
		builder.WriteString(letters[i])
	}
	*str = builder.String()

}
