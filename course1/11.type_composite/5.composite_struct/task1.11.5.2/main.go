package main

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	petname "github.com/dustinkirkland/golang-petname"
)

func main() {
	data := getAnimals()
	info := strings.Trim(preparePrint(data), "\n")
	fmt.Println(info)
}

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	animals := []Animal{}
	for i := 0; i < 3; i++ {
		age := gofakeit.Number(3, 40)
		name := petname.Generate(1, "")
		Type := gofakeit.AnimalType()
		animals = append(animals, Animal{Type: Type, Name: name, Age: age})
	}
	return animals
}

func preparePrint(animals []Animal) string {
	var sb strings.Builder
	for _, i := range animals {
		sb.WriteString(fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d\n", i.Type, i.Name, i.Age))
	}
	return sb.String()
}
