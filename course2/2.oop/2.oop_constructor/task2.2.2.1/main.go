package main

import (
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}
type OrderOption func(*Order)

func WithCustomerID(ID string) OrderOption {
	return func(o *Order) {
		o.CustomerID = ID
	}
}

func WithItems(items []string) OrderOption {
	return func(o *Order) {
		o.Items = items
	}
}

func WithOrderDate(date time.Time) OrderOption {
	return func(o *Order) {
		o.OrderDate = date
	}
}

func NewOrder(id int, options ...OrderOption) *Order {
	order := &Order{ID: id}

	for _, option := range options {
		option(order)
	}

	return order
}

func main() {
}
