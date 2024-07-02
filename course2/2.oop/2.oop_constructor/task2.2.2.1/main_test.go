package main

import (
	"testing"
	"time"
)

func TestCreatingOrderWithAllOptions(t *testing.T) {
	orderDate := time.Now()
	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(orderDate))

	if order.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", order.ID)
	}

	if order.CustomerID != "123" {
		t.Errorf("expected CustomerID to be '123', got %s", order.CustomerID)
	}

	if len(order.Items) != 2 || order.Items[0] != "item1" || order.Items[1] != "item2" {
		t.Errorf("expected Items to be ['item1', 'item2'], got %v", order.Items)
	}

	if !order.OrderDate.Equal(orderDate) {
		t.Errorf("expected OrderDate to be %v, got %v", orderDate, order.OrderDate)
	}
}

func TestCreatingOrderWithEmptyCustomerID(t *testing.T) {
	orderDate := time.Now()
	order := NewOrder(2,
		WithCustomerID(""),
		WithItems([]string{"item3", "item4"}),
		WithOrderDate(orderDate))

	if order.ID != 2 {
		t.Errorf("expected ID to be 2, got %d", order.ID)
	}

	if order.CustomerID != "" {
		t.Errorf("expected CustomerID to be empty, got %s", order.CustomerID)
	}

	if len(order.Items) != 2 || order.Items[0] != "item3" || order.Items[1] != "item4" {
		t.Errorf("expected Items to be ['item3', 'item4'], got %v", order.Items)
	}

	if !order.OrderDate.Equal(orderDate) {
		t.Errorf("expected OrderDate to be %v, got %v", orderDate, order.OrderDate)
	}
}
