package main

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit"
)

func main() {
	data := getUsers()
	info := strings.Trim(preparePrint(data), "\n")
	fmt.Println(info)
}

type User struct {
	Name string
	Age  int
}

func getUsers() []User {
	usersSlice := []User{}
	for i := 0; i < 10; i++ {
		name := gofakeit.Name()
		age := gofakeit.Number(18, 60)
		usersSlice = append(usersSlice, User{Name: name, Age: age})
	}
	return usersSlice
}

func preparePrint(data []User) string {
	var sb strings.Builder
	for _, i := range data {
		sb.WriteString(fmt.Sprintf("Имя: %s, Возраст: %d\n", i.Name, i.Age))
	}
	return sb.String()
}
