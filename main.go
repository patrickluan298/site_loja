package main

import (
	"net/http"
	"site_loja/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.HandleRoutes()
	http.ListenAndServe(":8000", nil)
}
