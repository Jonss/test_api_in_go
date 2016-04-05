package main

import (
			"encoding/json"
			"net/http"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	products := Products {
		Product{Name: "Batata", Url: "http://google.com.br", Store: "Extra"},
		Product{Name: "Ruffles", Url: "http://google.com.br", Store: "Pão de açucar"},
		Product{Name: "Pão de Batata", Url: "http://google.com.br", Store: "Extra"},
	}

	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/products", getProducts)
	http.ListenAndServe(":3000", nil)
}