// https://leetcode.com/problems/longest-common-prefix/description/
package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var sb strings.Builder

	for i := 0; i < len(strs[0]); i++ {
		letter := strs[0][i]

		for _, j := range strs {
			if len(j) <= i || j[i] != letter {
				return sb.String()
			}
		}

		sb.WriteByte(letter)
	}

	return sb.String()
}

func main() {
}
