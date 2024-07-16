package main

import (
	"encoding/json"

	"github.com/brianvoe/gofakeit/v6"
	jsoniter "github.com/json-iterator/go"
)

//go:generate easyjson -all $main.go
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

func EasyJSON(users []User) {
	panic("Write me")
}
func JSON(users []User) {
	for _, user := range users {
		json.Marshal(user)
	}
}

func JSONiter(users []User) {
	for _, user := range users {
		jsonIter := jsoniter.ConfigCompatibleWithStandardLibrary
		jsonIter.Marshal(user)
	}
}
func GenerateUSER(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			ID:       i,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, true, false, 14),
			Age:      gofakeit.Number(18, 100),
			Email:    gofakeit.Email(),
		}
	}
	return users
}

func main() {
}
