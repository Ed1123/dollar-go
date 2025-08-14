package prices

import (
	"encoding/json"
	"math"
	"math/rand"
	"net/http"
	"slices"
)

type ExchangeHouse struct {
	Name      string  `json:"exchange_name"`
	URL       string  `json:"url"`
	BuyPrice  float64 `json:"buy_price"`
	SellPrice float64 `json:"sell_price"`
	Logo      string  `json:"icon,omitempty"`
}

func GetPricesMock() ([]ExchangeHouse, error) {
	// Simulate a slow API call
	// time.Sleep(5 * time.Second)

	exchangeHouses := []ExchangeHouse{
		{
			Name:      "Exchange A",
			URL:       "https://example.com/exchange-a",
			Logo:      "https://images.unsplash.com/photo-1584044283481-9b18070011e1?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			BuyPrice:  math.Round((rand.Float64()+3)*1000) / 1000,
			SellPrice: math.Round((rand.Float64()+3)*1000) / 1000,
		},
		{
			Name:      "Exchange B",
			URL:       "https://example.com/exchange-b",
			Logo:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcREh1QdeZms_FPe7R2a_wGvpti_aNnJxQAs_g&s",
			BuyPrice:  math.Round((rand.Float64()+3)*1000) / 1000,
			SellPrice: math.Round((rand.Float64()+3)*1000) / 1000,
		},
		{
			Name:      "Exchange C",
			URL:       "https://example.com/exchange-c",
			Logo:      "https://images.unsplash.com/photo-1584044283481-9b18070011e1?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			BuyPrice:  math.Round((rand.Float64()+3)*1000) / 1000,
			SellPrice: math.Round((rand.Float64()+3)*1000) / 1000,
		},
		{
			Name:      "Exchange D",
			URL:       "https://example.com/exchange-d",
			Logo:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcREh1QdeZms_FPe7R2a_wGvpti_aNnJxQAs_g&s",
			BuyPrice:  math.Round((rand.Float64()+3)*1000) / 1000,
			SellPrice: math.Round((rand.Float64()+3)*1000) / 1000,
		},
	}

	return exchangeHouses, nil
}

func GetPrices() ([]ExchangeHouse, error) {
	response, err := http.Get("https://dollar-sol-api-812540214021.us-central1.run.app/rates")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, http.ErrNotSupported
	}
	houses := make([]ExchangeHouse, 0)
	err = json.NewDecoder(response.Body).Decode(&houses)
	if err != nil {
		return nil, err
	}
	if len(houses) == 0 {
		return nil, http.ErrContentLength
	}
	return houses, nil
}

type DollarPrice struct {
	Name  string
	URL   string
	Logo  string
	Price float64
}

type BestPrices struct {
	Buying  []DollarPrice
	Selling []DollarPrice
}

// BestExchangeHouses takes a slice of ExchangeHouse and returns
// the best places to buy and sell dollars.
// Buying is sorted so the lowest price is first.
// Selling is sorted so the highest price is first.
func BestExchangeHouses(houses []ExchangeHouse) BestPrices {
	bestBuying := make([]DollarPrice, 0)
	bestSelling := make([]DollarPrice, 0)

	for _, house := range houses {
		if house.BuyPrice > 0 {
			bestBuying = append(bestBuying, DollarPrice{
				Name:  house.Name,
				URL:   house.URL,
				Logo:  house.Logo,
				Price: house.BuyPrice,
			})
		}
		if house.SellPrice > 0 {
			bestSelling = append(bestSelling, DollarPrice{
				Name:  house.Name,
				URL:   house.URL,
				Logo:  house.Logo,
				Price: house.SellPrice,
			})
		}
	}

	// Sort bestBuying by Price ascending
	slices.SortFunc(bestBuying, func(a, b DollarPrice) int {
		if a.Price < b.Price {
			return -1
		} else if a.Price > b.Price {
			return 1
		}
		return 0
	})
	// Sort bestSelling by Price descending
	slices.SortFunc(bestSelling, func(a, b DollarPrice) int {
		if a.Price > b.Price {
			return -1
		} else if a.Price < b.Price {
			return 1
		}
		return 0
	})

	return BestPrices{
		Buying:  bestBuying,
		Selling: bestSelling,
	}
}
