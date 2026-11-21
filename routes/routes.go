package routes

import (
	"net/http"

	"pokedex/controllers"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/pokemon", controllers.List)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	return mux
}
