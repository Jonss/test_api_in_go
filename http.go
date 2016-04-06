package main

import (
		"net/http"
		"encoding/json"	
)

func handler(w http.ResponseWriter, r *http.Request){	


	products := getProducts()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(products)
}


func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":3000", nil)

}