package main

import (
	"strings"
	"unicode"
)

func main() {
}

func filterWords(text string, censorMap map[string]string) string {
	if text == "" {
		return ""
	}

	sentences := splitSentences(text)

	words := strings.Fields(sentences[0])
	passedWords := []string{}
	uniqueWords := make(map[string]bool)

	for _, i := range words {
		value, ok := censorMap[strings.ToLower(i)]
		if ok {
			i = checkUpper(i, value)
		}
		_, ok = uniqueWords[strings.ToLower(i)]
		if !ok {
			passedWords = append(passedWords, i)
			uniqueWords[strings.ToLower(i)] = true
		}
	}

	answer := strings.Join(passedWords, " ") + "!"
	if len(sentences) > 1 {
		for i := 1; i < len(sentences); i++ {
			answer = answer + " " + (filterWords(WordsToSentence(strings.Fields(sentences[i])), censorMap))
		}
	}

	return strings.Trim(answer, " ")
}

func WordsToSentence(words []string) string {
	filtered := make([]string, 0, len(words))

	for _, word := range words {
		if word != "" {
			filtered = append(filtered, word)
		}
	}

	return strings.ReplaceAll(strings.Join(filtered, " ")+"!", "!!", "!")
}

func checkUpper(old, new string) string {
	if len(old) == 0 || len(new) == 0 {
		return new
	}

	chars := []rune(old)

	if unicode.IsUpper(chars[0]) {
		runes := []rune(new)
		new = string(append([]rune{unicode.ToUpper(runes[0])}, runes[1:]...))
	}

	return new
}

func splitSentences(message string) []string {
	originSentences := strings.Split(message, "!")
	if len(originSentences) == 2 && originSentences[1] == "" {
		return []string{originSentences[0]}
	}
	var orphan string
	var sentences []string
	for i, sentence := range originSentences {
		words := strings.Split(sentence, " ")
		if len(words) == 1 {
			if len(orphan) > 0 {
				orphan += " "
			}

			orphan += words[0] + "!"
			continue
		}
		if orphan != "" {
			originSentences[i] = strings.Join([]string{orphan, " ", sentence}, " ") + "!"
			orphan = ""
		}
		sentences = append(sentences, originSentences[i])
	}
	return sentences
}
