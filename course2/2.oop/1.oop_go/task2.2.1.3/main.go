package main

import (
	"fmt"
)

type DineInOrder struct {
	orderDetails map[string]int
}

type TakeAwayOrder struct {
	orderDetails map[string]int
}

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

func (order *DineInOrder) AddItem(item string, quantity int) error {
	order.orderDetails[item] += quantity
	return nil
}

func (order *TakeAwayOrder) AddItem(item string, quantity int) error {
	order.orderDetails[item] += quantity
	return nil
}

func (order *DineInOrder) RemoveItem(item string) error {
	if order.orderDetails[item] < 1 {
		return fmt.Errorf("not enough %s in order", item)
	}
	order.orderDetails[item]--
	return nil
}

func (order *TakeAwayOrder) RemoveItem(item string) error {
	if order.orderDetails[item] < 1 {
		return fmt.Errorf("not enough %s in order", item)
	}
	order.orderDetails[item]--
	return nil
}

func (order *TakeAwayOrder) GetOrderDetails() map[string]int {
	return order.orderDetails
}

func (order *DineInOrder) GetOrderDetails() map[string]int {
	return order.orderDetails
}

func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	ManageOrder(dineIn)
	ManageOrder(takeAway)
}
