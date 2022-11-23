package models

import (
	"store/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func SearchProducts() []Product {
	db := db.ConectDatabase()

	selectQueryAll, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectQueryAll.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectQueryAll.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)

	}

	defer db.Close()
	return products

}

func ConectDatabase() {
	panic("unimplemented")
}

func InsertProduct(name string, description string, price float64, amount int) {
	db := db.ConectDatabase()

	insertQuery, err := db.Prepare("INSERT INTO products(name, description, price, amount) VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(name, description, price, amount)

	defer db.Close()
}