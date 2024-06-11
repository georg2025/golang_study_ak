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
	var answer strings.Builder
	var sb strings.Builder
	words := strings.Fields(sentences[0])

	uniqueWords := make(map[string]bool)

	for _, i := range words {
		value, ok := censorMap[strings.ToLower(i)]
		if ok {
			i = checkUpper(i, value)
		}
		_, ok = uniqueWords[strings.ToLower(i)]
		if !ok {
			sb.WriteString(i)
			sb.WriteString(" ")
			uniqueWords[strings.ToLower(i)] = true
		}
	}

	answer.WriteString(strings.Trim(sb.String(), " "))
	answer.WriteString("! ")
	if len(sentences) > 1 {
		for i := 1; i < len(sentences); i++ {
			answer.WriteString(filterWords(WordsToSentence(strings.Fields(sentences[i])), censorMap))
		}
	}

	return answer.String()
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
