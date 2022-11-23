package main

import (
	"net/http"
	"store/routes"

	_ "github.com/lib/pq"
)

func main() {

	routes.Routes()
	http.ListenAndServe(":8080", nil)

}
