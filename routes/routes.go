package routes

import (
	"net/http"
	"store/controllers"
)

func LoadingRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
