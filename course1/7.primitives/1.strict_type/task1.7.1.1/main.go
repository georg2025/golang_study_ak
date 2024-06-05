package main

import "fmt"

func main() {
	var name string
	var age int
	var city string
	fmt.Print("Введите ваше имя: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		panic(err)
	}

	fmt.Print("Введите ваш возраст: ")
	_, err = fmt.Scanln(&age)
	if err != nil {
		panic(err)
	}

	fmt.Print("Введите ваш город: ")
	_, err = fmt.Scanln(&city)
	if err != nil {
		panic(err)
	}

	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Город:", city)
}
