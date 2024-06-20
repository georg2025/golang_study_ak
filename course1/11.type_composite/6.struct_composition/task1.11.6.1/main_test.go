package main

import (
	"testing"
)

func TestOrder(t *testing.T) {
	order := Order{Dishes: []Dish{}, Total: 0.0}
	order.AddDish(Dish{Name: "tea", Price: 10.2})
	if len(order.Dishes) != 1 {
		t.Errorf("Dishes dont add")
	}

	order.CaltulateTotal()
	result := order.Total
	expected := 10.2
	if result != expected {
		t.Errorf("Got %f, expected %f", result, expected)
	}

	order.RemoveDish(Dish{Name: "coconut", Price: 90.2})
	if len(order.Dishes) != 1 {
		t.Errorf("It removes wrong items")
	}

	order.RemoveDish(Dish{Name: "tea", Price: 10.2})
	if len(order.Dishes) != 0 {
		t.Errorf("It doesn't remove")
	}
}
