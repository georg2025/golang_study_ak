package main

import (
	"regexp"
	"strings"
)

func main() {
}

type Ad struct {
	Title       string
	Discription string
}

func censorAds(ads []Ad, censor map[string]string) []Ad {
	if len(ads) == 0 {
		return []Ad{}
	}
	re := regexp.MustCompile(makeStringForRegex(censor))
	replacer := func(match string) string { return censor[strings.ToLower(match)] }

	answer := []Ad{}
	for _, i := range ads {
		outputTitle := re.ReplaceAllStringFunc(i.Title, replacer)
		outputText := re.ReplaceAllStringFunc(i.Discription, replacer)
		i.Title = outputTitle
		i.Discription = outputText
		answer = append(answer, i)
	}
	return answer
}

func makeStringForRegex(censor map[string]string) string {
	var sb strings.Builder
	for i := range censor {
		sb.WriteString("(?i)")
		sb.WriteString(i)
		sb.WriteString("|")
	}
	answer := sb.String()
	answer = strings.Trim(answer, "|")
	return answer
}
