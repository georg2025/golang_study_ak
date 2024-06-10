package main

import "reflect"

func main() {
}

type Dish struct {
	Name  string
	Price float64
}

type Order struct {
	Dishes []Dish
	Total  float64
}

func (order *Order) AddDish(dish Dish) {
	order.Dishes = append(order.Dishes, dish)
}

func (order *Order) RemoveDish(dish Dish) {
	number := -1
	for j, i := range order.Dishes {
		if reflect.DeepEqual(i, dish) {
			number = j
			break
		}
	}
	if number >= 0 {
		newOrder := order.Dishes[:number]
		newOrder = append(newOrder, order.Dishes[number+1:]...)
		order.Dishes = newOrder
	}
}

func (order *Order) CaltulateTotal() {
	total := 0.0
	for _, i := range order.Dishes {
		total += i.Price
	}
	order.Total = total
}
