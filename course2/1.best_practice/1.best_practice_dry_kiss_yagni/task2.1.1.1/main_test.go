package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	product := &Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	product.ProductID = ProductCocaCola
	product.ProductID = ProductPepsi
	product.ProductID = ProductSprite

	statProfit := NewStatisticProfit(
		WithAverageProfit,
		WithAverageProfitPercent,
		WithCurrentProfit,
		WithDiffenceProfit,
		WithAllData,
	).(*StatisticProfit)

	statProfit.SetProduct(product)
	expectedData := make([]float64, 0, 4)

	result := statProfit.product.ProductID
	if result != 2 {
		t.Errorf("Expected %d, got %d", 2, result)
	}

	resultFloat := statProfit.getAverageProfit()
	if resultFloat != 5.0 {
		t.Errorf("Expected %v, got %v", 5.0, resultFloat)
	}

	expectedData = append(expectedData, 5.0)

	resultFloat = statProfit.getAverageProfitPercent()
	expected := statProfit.getAverageProfit() / statProfit.Average(statProfit.product.Buys) * 100
	if resultFloat != expected {
		t.Errorf("Expected %v, got %v", expected, resultFloat)
	}

	expectedData = append(expectedData, expected)

	resultFloat = statProfit.getCurrentProfit()
	if resultFloat != 3.5 {
		t.Errorf("Expected %v, got %v", 3.5, resultFloat)
	}

	resultFloat = statProfit.getDifferenceProfit()
	if resultFloat != 15.0 {
		t.Errorf("Expected %v, got %v", 15.0, resultFloat)
	}

	resultData := statProfit.getAllData()
	expectedData = append(expectedData, 3.5, 15.0)
	if !reflect.DeepEqual(resultData, expectedData) {
		t.Errorf("Expected %v, got %v", expectedData, resultData)
	}
}
