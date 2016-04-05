package main


type Product struct {
	Name string `json:"name"`
	Url string `json:"url" `
	Store string `json:"store" `
}

type Products []Product