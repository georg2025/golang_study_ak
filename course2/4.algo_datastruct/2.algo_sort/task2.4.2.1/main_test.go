package main

import (
	"sort"
	"testing"
	"time"
)

func TestSorting(t *testing.T) {
	products := []Product{
		Product{
			Name:      "product1",
			Price:     30.0,
			CreatedAt: time.Now(),
			Count:     1,
		},
		Product{
			Name:      "product2",
			Price:     20.0,
			CreatedAt: time.Now().Add(-time.Hour * 2),
			Count:     3,
		},
		Product{
			Name:      "product3",
			Price:     10.0,
			CreatedAt: time.Now().Add(-time.Hour),
			Count:     2,
		},
	}

	sort.Sort(ByPrice(products))
	if products[0].Name != "product3" {
		t.Errorf("expected name product3, got %v", products[0].Name)
	}

	sort.Sort(ByCreatedAt(products))
	if products[0].Name != "product2" {
		t.Errorf("expected name product2, got %v", products[0].Name)
	}

	sort.Sort(ByCount(products))
	if products[0].Name != "product1" {
		t.Errorf("expected name product1, got %v", products[0].Name)
	}
}

func TestGenerateProducts(t *testing.T) {
	n := 5
	products := generateProducts(n)

	if len(products) != n {
		t.Errorf("Expected %d products, but got %d", n, len(products))
	}
}
