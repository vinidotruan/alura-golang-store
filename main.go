package main

import (
	"net/http"
	"vinidotruan/go-store/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
