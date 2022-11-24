package controllers

import (
	"html/template"
	"log"
	"net/http"
	"store/models"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	AllProducts := models.SearchProducts()

	templates.ExecuteTemplate(w, "Index", AllProducts)

}
func New(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "New", nil)

}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price to float64: ", err)
		}
		amountConverted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error converting amount to int: ", err)
		}

		models.InsertProduct(name, description, priceConverted, amountConverted)
	}

	http.Redirect(w, r, "/", 301)

}
func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")

	models.DeleteProduct(idProduct)

	http.Redirect(w, r, "/", 301)

}
func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")

	product := models.EditProduct(idProduct)

	templates.ExecuteTemplate(w, "Edit", product)

}
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price to float64: ", err)
		}
		amountConverted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error converting amount to int: ", err)
		}
		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error converting id to int: ", err)
		}

		models.UpdateProduct(idConverted, name, description, priceConverted, amountConverted)
	}
	http.Redirect(w, r, "/", 301)
}
