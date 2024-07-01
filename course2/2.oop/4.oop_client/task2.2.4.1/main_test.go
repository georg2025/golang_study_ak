package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTickerReturnsValidData(t *testing.T) {
	// Mock server to return a valid Ticker response
	mockResponse := `{
    "BTC_USD": {
      "buy_price": "50000",
      "sell_price": "51000",
      "last_trade": "50500",
      "high": "52000",
      "low": "48000",
      "avg": "50000",
      "vol": "1000",
      "vol_curr": "50000000",
      "updated": 1617184800
    }
  }`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Initialize Exmo with the mock server URL
	exmo := NewExmo(WithURL(server.URL))

	// Call GetTicker
	ticker, err := exmo.GetTicker()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Validate the response
	expected := Ticker{
		"BTC_USD": {
			BuyPrice:  "50000",
			SellPrice: "51000",
			LastTrade: "50500",
			High:      "52000",
			Low:       "48000",
			Avg:       "50000",
			Vol:       "1000",
			VolCurr:   "50000000",
			Updated:   1617184800,
		},
	}

	if !jsonEqual(ticker, expected) {
		t.Errorf("expected %v, got %v", expected, ticker)
	}
}

func jsonEqual(a, b interface{}) bool {
	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)
	return string(aJSON) == string(bJSON)
}
