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

func ConectDatabase() {
	panic("unimplemented")
}

func SearchProducts() []Product {
	db := db.ConectDatabase()

	selectQueryAll, err := db.Query("select * from products order by id asc")
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

func InsertProduct(name string, description string, price float64, amount int) {
	db := db.ConectDatabase()

	insertQuery, err := db.Prepare("INSERT INTO products(name, description, price, amount) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(name, description, price, amount)

	defer db.Close()
}

func DeleteProduct(idProduct string) {
	db := db.ConectDatabase()

	deleteQuery, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteQuery.Exec(idProduct)

	defer db.Close()

}

func EditProduct(id string) Product {
	db := db.ConectDatabase()

	productData, err := db.Query("SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	productAtt := Product{}
	for productData.Next() {
		var id, amount int
		var name, description string
		var price float64
		err = productData.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		productAtt.Id = id
		productAtt.Name = name
		productAtt.Description = description
		productAtt.Price = price
		productAtt.Amount = amount

	}
	defer db.Close()
	return productAtt

}

func UpdateProduct(id int, name string, description string, price float64, amount int) {
	db := db.ConectDatabase()

	updateQuery, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, amount=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateQuery.Exec(name, description, price, amount, id)

	defer db.Close()

}
