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

func Delete(idProduct string) {
	connection := db.DatabaseCon()

	delete, err := connection.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(idProduct)
	defer connection.Close()
}

func GetById(id string) Product {
	connection := db.DatabaseCon()
	findProductQuery, err := connection.Query("SELECT * FROM products WHERE id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	for findProductQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = findProductQuery.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Price = price
		product.Description = description
		product.Quantity = quantity

	}
	connection.Close()
	return product
}

func Update(id int, name string, description string, quantity int, price float64) {
	connection := db.DatabaseCon()

	update, err := connection.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(name, description, price, quantity, id)
	defer connection.Close()
}
