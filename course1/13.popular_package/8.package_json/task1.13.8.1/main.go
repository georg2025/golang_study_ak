package main

import (
	"encoding/json"
	"strings"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Text string `json:"text"`
}

func main() {
}

func getJSON(data []User) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	var sb strings.Builder
	for _, i := range data {
		jsonData, err := json.Marshal(i)
		if err != nil {
			return "", err
		}
		sb.WriteString(string(jsonData))
	}
	return sb.String(), nil
}
