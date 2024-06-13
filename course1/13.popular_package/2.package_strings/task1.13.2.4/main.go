package main

import (
	"math/rand"
	"strings"
)

func main() {
}

func GenerateActivationKey() string {
	letters := []string{}
	for i := 0; i < 4; i++ {
		letters = append(letters, string(GenerateRandomChars(100)))
	}

	return strings.Join(letters, "-")
}

func GenerateRandomChars(seed int64) []rune {
	letters := []rune("abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	answer := []rune{}
	rand.Seed(seed)
	for i := 0; i < 4; i++ {
		randomNumber := rand.Intn(len(letters))
		answer = append(answer, letters[randomNumber])
	}
	return answer
}
