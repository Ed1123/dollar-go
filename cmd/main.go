package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/Ed1123/dollar-go/prices"
)

var IsTest bool

func main() {
	IsTest = strings.ToLower(os.Getenv("ENV")) == "dev"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/prices", pricesHandler)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}
	slog.Info("Server started", "url", "http://localhost:"+port, "test", IsTest)
	http.ListenAndServe(":"+port, nil)
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
	getPrices := prices.GetPrices
	if IsTest {
		slog.Info("Running in test mode, using mock prices")
		getPrices = prices.GetPricesMock
	}
	data, err := getPrices()
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
