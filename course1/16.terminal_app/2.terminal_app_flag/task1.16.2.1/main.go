package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var num int
	flag.IntVar(&num, "n", 0, "drawing deep")
	flag.Parse()
	path := os.Args[3:]
	printTree(path[0], "", true, num)
}

func printTree(path string, prefix string, isLast bool, depth int) {
	pathSlice := strings.Split(path, "/")
	firstElemToTake := len(pathSlice) - depth

	if firstElemToTake < 0 {
		firstElemToTake = 0
	}

	pathSlice = pathSlice[firstElemToTake:]

	if len(pathSlice) == 0 {
		return
	}

	if len(pathSlice) == 1 {
		fmt.Println(pathSlice[0])
		return
	}

	var sb strings.Builder

	firstString := fmt.Sprintf("%s\n  |-----", pathSlice[0])
	secondString := strings.Join(pathSlice[1:], "-----")
	sb.WriteString(prefix)
	sb.WriteString(firstString)
	sb.WriteString(secondString)

	if !isLast {
		sb.WriteString("...")
	}

	fmt.Println(sb.String())
}
