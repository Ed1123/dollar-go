package prices

import (
	"reflect"
	"testing"
)

func TestBestExchangeHouses_SortsCorrectly(t *testing.T) {
	houses := []ExchangeHouse{
		{
			Name:      "A",
			URL:       "urlA",
			Logo:      "logoA",
			BuyPrice:  100,
			SellPrice: 110,
		},
		{
			Name:      "B",
			URL:       "urlB",
			Logo:      "logoB",
			BuyPrice:  90,
			SellPrice: 120,
		},
		{
			Name:      "C",
			URL:       "urlC",
			Logo:      "logoC",
			BuyPrice:  95,
			SellPrice: 105,
		},
	}

	wantBuying := []DollarPrice{
		{Name: "B", URL: "urlB", Logo: "logoB", Price: 90},
		{Name: "C", URL: "urlC", Logo: "logoC", Price: 95},
		{Name: "A", URL: "urlA", Logo: "logoA", Price: 100},
	}
	wantSelling := []DollarPrice{
		{Name: "B", URL: "urlB", Logo: "logoB", Price: 120},
		{Name: "A", URL: "urlA", Logo: "logoA", Price: 110},
		{Name: "C", URL: "urlC", Logo: "logoC", Price: 105},
	}

	got := BestExchangeHouses(houses)

	if !reflect.DeepEqual(got.Buying, wantBuying) {
		t.Errorf("Buying not sorted as expected.\nGot:  %#v\nWant: %#v", got.Buying, wantBuying)
	}
	if !reflect.DeepEqual(got.Selling, wantSelling) {
		t.Errorf("Selling not sorted as expected.\nGot:  %#v\nWant: %#v", got.Selling, wantSelling)
	}
}

func TestBestExchangeHouses_ZeroOrNegativePrices(t *testing.T) {
	houses := []ExchangeHouse{
		{Name: "A", URL: "urlA", Logo: "logoA", BuyPrice: 0, SellPrice: 100},
		{Name: "B", URL: "urlB", Logo: "logoB", BuyPrice: -10, SellPrice: -20},
		{Name: "C", URL: "urlC", Logo: "logoC", BuyPrice: 50, SellPrice: 0},
	}

	got := BestExchangeHouses(houses)

	wantBuying := []DollarPrice{
		{Name: "C", URL: "urlC", Logo: "logoC", Price: 50},
	}
	wantSelling := []DollarPrice{
		{Name: "A", URL: "urlA", Logo: "logoA", Price: 100},
	}

	if !reflect.DeepEqual(got.Buying, wantBuying) {
		t.Errorf("Buying with zero/negative prices failed. Got: %#v, Want: %#v", got.Buying, wantBuying)
	}
	if !reflect.DeepEqual(got.Selling, wantSelling) {
		t.Errorf("Selling with zero/negative prices failed. Got: %#v, Want: %#v", got.Selling, wantSelling)
	}
}

func TestBestExchangeHouses_EmptyInput(t *testing.T) {
	got := BestExchangeHouses([]ExchangeHouse{})
	if len(got.Buying) != 0 {
		t.Errorf("Expected empty Buying, got: %#v", got.Buying)
	}
	if len(got.Selling) != 0 {
		t.Errorf("Expected empty Selling, got: %#v", got.Selling)
	}
}
