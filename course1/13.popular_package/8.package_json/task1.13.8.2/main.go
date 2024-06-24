package main

import (
	"encoding/json"
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

func getUsersFromJSON(data []byte) ([]User, error) {
	if len(data) == 0 {
		return []User{}, nil
	}

	users := []User{}
	err := json.Unmarshal(data, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}
