package main

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/Ed1123/dollar-go/prices"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/prices", pricesHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func pricesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := prices.GetPricesMock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Info("Fetched prices", "count", len(data))
	bestPrices := prices.BestExchangeHouses(data)
	tmpl, err := template.ParseFiles("templates/prices_list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		slog.Error("Failed to parse template", "error", err)
		return
	}
	err = tmpl.Execute(w, bestPrices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		slog.Error("Failed to execute template", "error", err)
		return
	}
}
