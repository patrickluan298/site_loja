package main

import (
	"net/http"
	"site_loja/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.HandleRequest()
	http.ListenAndServe(":8000", nil)
}
