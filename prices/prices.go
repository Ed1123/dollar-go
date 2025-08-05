package prices

import (
	"math/rand"
)

type ExchangeHouse struct {
	Name      string  `json:"exchange_name"`
	URL       string  `json:"url"`
	BuyPrice  float64 `json:"buy_price"`
	SellPrice float64 `json:"sell_price"`
	Logo      string  `json:"logo,omitempty"`
}

func GetPricesMock() ([]ExchangeHouse, error) {
	// Simulate a slow API call
	// time.Sleep(5 * time.Second)

	exchangeHouses := []ExchangeHouse{
		{
			Name:      "Exchange A",
			URL:       "https://example.com/exchange-a",
			Logo:      "https://images.unsplash.com/photo-1584044283481-9b18070011e1?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			BuyPrice:  rand.Float64() * 100,
			SellPrice: rand.Float64() * 100,
		},
		{
			Name:      "Exchange B",
			URL:       "https://example.com/exchange-b",
			Logo:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcREh1QdeZms_FPe7R2a_wGvpti_aNnJxQAs_g&s",
			BuyPrice:  rand.Float64() * 100,
			SellPrice: rand.Float64() * 100,
		},
	}

	return exchangeHouses, nil
}
