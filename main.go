package main

import (
	"net/http"
	"store/routes"
)

func main() {
	routes.LoadingRoutes()
	http.ListenAndServe(":8080", nil)
}
