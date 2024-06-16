package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.NewBufferString("Hello, World!")
	b := make([]byte, 13)
	r := getReader(buffer)
	r.Read(b)
	fmt.Println(string(b))
}

// This is our function. We get buffer and use it as source to make a new Reader (we can make it, cause)
// bytes.Buffer implements Reader interface.
func getReader(b *bytes.Buffer) *bufio.Reader {
	return bufio.NewReader(b)
}
