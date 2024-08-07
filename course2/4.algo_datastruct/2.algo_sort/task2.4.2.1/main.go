package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Product struct {
	Name      string
	Price     float64
	CreatedAt time.Time
	Count     int
}

func (p Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %f, Count: %v", p.Name, p.Price, p.Count)
}

func generateProducts(n int) []Product {
	gofakeit.Seed(time.Now().UnixNano())
	products := make([]Product, n)
	for i := range products {
		products[i] = Product{
			Name:      gofakeit.Word(),
			Price:     gofakeit.Price(1.0, 100.0),
			CreatedAt: gofakeit.Date(),
			Count:     gofakeit.Number(1, 100),
		}
	}
	return products
}

type ByPrice []Product

func (price ByPrice) Len() int           { return len(price) }
func (price ByPrice) Less(i, j int) bool { return price[i].Price < price[j].Price }
func (price ByPrice) Swap(i, j int)      { price[i], price[j] = price[j], price[i] }

type ByCreatedAt []Product

func (created ByCreatedAt) Len() int { return len(created) }
func (created ByCreatedAt) Less(i, j int) bool {
	return created[i].CreatedAt.Before(created[j].CreatedAt)
}
func (created ByCreatedAt) Swap(i, j int) { created[i], created[j] = created[j], created[i] }

type ByCount []Product

func (count ByCount) Len() int           { return len(count) }
func (count ByCount) Less(i, j int) bool { return float64(count[i].Count) < float64(count[j].Count) }
func (count ByCount) Swap(i, j int)      { count[i], count[j] = count[j], count[i] }

func main() {
}
