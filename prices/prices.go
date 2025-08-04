package prices

import (
	"math/rand"
	"time"
)

type ExchangeHouse struct {
	Name   string       `json:"name"`
	URL    string       `json:"url"`
	Icon   string       `json:"icon"`
	Prices DollarPrices `json:"prices"`
}

type DollarPrices struct {
	Selling float64 `json:"selling"`
	Buying  float64 `json:"buying"`
}

func GetPricesMock() ([]ExchangeHouse, error) {
	// Simulate a slow API call
	time.Sleep(5 * time.Second)

	exchangeHouses := []ExchangeHouse{
		{
			Name: "Exchange A",
			URL:  "https://example.com/exchange-a",
			Icon: "🏦",
			Prices: DollarPrices{
				Selling: rand.Float64() * 100,
				Buying:  rand.Float64() * 100,
			},
		},
		{
			Name: "Exchange B",
			URL:  "https://example.com/exchange-b",
			Icon: "🏧",
			Prices: DollarPrices{
				Selling: rand.Float64() * 100,
				Buying:  rand.Float64() * 100,
			},
		},
	}

	return exchangeHouses, nil
}
