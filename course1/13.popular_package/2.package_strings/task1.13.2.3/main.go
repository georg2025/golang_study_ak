package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println(GenerateRandomString(10))
}

func GenerateRandomString(length int) string {
	var sb strings.Builder
	letters := GenerateRandomChars(length)
	for _, i := range letters {
		sb.WriteRune(i)
	}
	return sb.String()
}

func GenerateRandomChars(number int) []rune {
	letters := []rune("abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	answer := []rune{}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < number; i++ {
		randomNumber := rand.Intn(len(letters))
		answer = append(answer, letters[randomNumber])
	}
	return answer
}
