package main

import (
	"fmt"
)

func main() {
	t := 7
	m := 14
	y := Dereference(&t)
	fmt.Println(y)
	f := Sum(&t, &m)
	fmt.Println(f)

}

func Dereference(n *int) int {
	return *n
}

func Sum(a, b *int) int {
	return *a + *b
}
