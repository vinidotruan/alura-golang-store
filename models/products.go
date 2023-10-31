package models

import db "vinidotruan/go-store/db"

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Quantity    int
}

func GetAll() []Product {
	connection := db.DatabaseCon()

	selectAllProducts, err := connection.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Price = price
		p.Description = description
		p.Quantity = quantity

		products = append(products, p)
	}
	defer connection.Close()
	return products
}

func CreateNew(name string, description string, quantity int, price float64) {
	connection := db.DatabaseCon()

	insert, err := connection.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)
	defer connection.Close()
}
