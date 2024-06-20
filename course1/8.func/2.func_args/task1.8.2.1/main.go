package main

import "fmt"

func main() {
	fmt.Println(UserInfo("Ivan", "Ivanovo", "88005553535", 33, 75))
}

func UserInfo(name, city, phone string, age, weight int) string {
	return fmt.Sprintf("Имя: %s, Город: %s, Телефон: %s, Возраст: %d, Вес: %d",
		name, city, phone, age, weight)
}
