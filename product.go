package main

	type Product struct {
		Position int `json:"position"`
		Name string `json:"name"`
		Store string `json:"store"`
		Url string  `json:"url"`
	}

	type Products []Product
