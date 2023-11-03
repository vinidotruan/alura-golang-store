package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"vinidotruan/go-store/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAll()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, errPrice := strconv.ParseFloat(r.FormValue("price"), 64)
		quantity, errQuantity := strconv.Atoi(r.FormValue("quantity"))

		if errPrice != nil {
			log.Println("Error converting price to float64", errPrice)
		}
		if errQuantity != nil {
			log.Println("Error converting quantity to integer", errQuantity)
		}

		models.CreateNew(name, description, quantity, price)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.Delete(idProduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.GetById(idProduct)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, errId := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, errPrice := strconv.ParseFloat(r.FormValue("price"), 64)
		quantity, errQuantity := strconv.Atoi(r.FormValue("quantity"))

		if errId != nil {
			log.Println("Error converting id to integer", errId)
		}
		if errPrice != nil {
			log.Println("Error converting price to float64", errPrice)
		}
		if errQuantity != nil {
			log.Println("Error converting quantity to integer", errQuantity)
		}

		models.Update(id, name, description, quantity, price)
		http.Redirect(w, r, "/", 301)
	}
}
