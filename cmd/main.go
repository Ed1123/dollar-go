package main

import (
	"htmx-go-app/prices"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/data", dataHandler)

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

func dataHandler(w http.ResponseWriter, r *http.Request) {
	data, err := prices.GetPricesMock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/partials/data.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
