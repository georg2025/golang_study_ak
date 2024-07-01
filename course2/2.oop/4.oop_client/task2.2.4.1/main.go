package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type OrderBookPair struct {
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BidQuantity string     `json:"bid_quantity"`
	BidAmount   string     `json:"bid_amount"`
	BidTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}

type TickerValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

type Pair struct {
	TradeID  int64  `json:"trade_id"`
	Date     int64  `json:"date"`
	Type     string `json:"type"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
}

type Currencies []string
type OrderBook map[string]OrderBookPair
type Ticker map[string]TickerValue
type Trades map[string][]Pair
type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo(opts ...func(exmo *Exmo)) *Exmo {
	exmo := &Exmo{client: &http.Client{}, url: "https://api.exmo.com/v1.1"}

	for _, option := range opts {
		option(exmo)
	}

	return exmo
}

func WithClient(client *http.Client) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.client = client
	}
}

func WithURL(url string) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.url = url
	}
}

func GetRespBody(e *Exmo, request string) ([]byte, error) {
	resp, err := e.client.Get(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func UnmarshalPair(data []byte) ([]Pair, error) {
	var r []Pair
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalOrderBookPair(data []byte) (OrderBookPair, error) {
	var r OrderBookPair
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalTicker(data []byte) (Ticker, error) {
	var r Ticker
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalCurrencies(data []byte) (Currencies, error) {
	var r Currencies
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalCandles(data []byte) (CandlesHistory, error) {
	var r CandlesHistory
	err := json.Unmarshal(data, &r)
	return r, err
}

func (e *Exmo) GetTicker() (Ticker, error) {
	request := e.url + ticker
	body, err := GetRespBody(e, request)

	if err != nil {
		return nil, err
	}

	info, _ := UnmarshalTicker([]byte(string(body)))
	return info, nil
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	requestSample := e.url + trades + "/?pair="
	book := Trades{}

	for _, pair := range pairs {
		request := requestSample + pair
		body, err := GetRespBody(e, request)

		if err != nil {
			return nil, err
		}

		data := make(map[string]json.RawMessage)
		err = json.Unmarshal(body, &data)

		if err != nil {
			return nil, err
		}

		var info []Pair

		for i, j := range data {
			info, err = UnmarshalPair(j)
			book[i] = info

			if err != nil {
				return nil, err
			}
		}
	}

	return book, nil
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	book := OrderBook{}

	for _, pair := range pairs {
		request := e.url + orderBook + "/?pair=" + pair + "&limit=" + strconv.Itoa(limit)

		body, err := GetRespBody(e, request)

		if err != nil {
			return nil, err
		}

		data := make(map[string]json.RawMessage)
		err = json.Unmarshal(body, &data)

		if err != nil {
			return nil, err
		}

		var info OrderBookPair

		for i, j := range data {
			info, err = UnmarshalOrderBookPair(j)

			if err != nil {
				return nil, err
			}

			book[i] = info
		}

	}

	return book, nil
}

func (e *Exmo) GetCurrencies() (Currencies, error) {
	request := e.url + currency
	body, err := GetRespBody(e, request)

	if err != nil {
		return nil, err
	}

	var info Currencies
	info, err = UnmarshalCurrencies(body)

	if err != nil {
		return nil, err
	}

	return info, nil
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	request := e.url + candlesHistory + "?symbol=" + pair + "&resolution=" + strconv.Itoa(limit) + "&from=" +
		strconv.FormatInt(start.Unix(), 10) + "&to=" + strconv.FormatInt(end.Unix(), 10)
	body, err := GetRespBody(e, request)
	var history CandlesHistory

	if err != nil {
		return history, err
	}

	history, err = UnmarshalCandles(body)

	if err != nil {
		return history, err
	}

	return history, nil
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	closedPrices := []float64{}
	history, err := e.GetCandlesHistory(pair, limit, start, end)

	if err != nil {
		return closedPrices, err
	}

	for _, candle := range history.Candles {
		closedPrices = append(closedPrices, candle.C)
	}

	return closedPrices, nil
}

// Другие методы...
func main() {
}
