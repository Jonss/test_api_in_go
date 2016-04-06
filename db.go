package main

import (
	 "database/sql"
	 "log"
	 _ "github.com/go-sql-driver/mysql"
	 "time"
		"fmt"
	 )


func getProducts() Products {

		begin := time.Now()
		products := make([]Product, 0)
		db, err := sql.Open("mysql", "root@tcp(:3306)/mallu_db")
			if( err != nil) {
				log.Fatal(err)
		}
		defer db.Close()

		rows, err := db.Query("SELECT * FROM products")
		if( err != nil) {
			log.Fatal(err)
		}	

		for rows.Next() {
			var pos int
			var name string
			var url string
			var store string
			err = rows.Scan(&pos, &name, &url, &store) 
			if err != nil {
				fmt.Println("yey deu ruim")
			}
			products = append(products, Product{pos, name, url, store})
		}
		
		db.Close()

		fmt.Println(time.Since(begin))
		return products

}

