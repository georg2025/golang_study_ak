package main

import (
	"math/rand"
	"strings"
	"time"
)

func main() {
}

// 255 - test mode. Everything else - normal mode
func GenerateActivationKey(mode uint8) string {
	letters := []string{}
	for i := 0; i < 4; i++ {
		if mode == 255 {
			letters = append(letters, string(GenerateRandomChars(100)))
		} else {
			letters = append(letters, string(GenerateRandomChars(time.Now().UTC().UnixNano())))
		}
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
