package main

import (
	"reflect"
	"unsafe"
)

func main() {
}

func getStringHeader(s string) reflect.StringHeader {
	return reflect.StringHeader{
		Data: uintptr((unsafe.Pointer(&s))),
		Len:  len(s),
	}
}
