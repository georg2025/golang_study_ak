package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(UserInfo("Ivan", 33, "Ivanovo", "Moscow", "Tver"))
}

func UserInfo(name string, age int, cities ...string) string {
	builder := &strings.Builder{}
	builder.WriteString("Имя: ")
	builder.WriteString(name)
	builder.WriteString(", возраст: ")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString(", города: ")
	if len(cities) == 0 {
		builder.WriteString(" нет таких")
		return builder.String()
	}
	for _, city := range cities {
		builder.WriteString(city)
		if city != cities[len(cities)-1] {
			builder.WriteString(", ")
		}
	}

	return builder.String()
}
