package controllers

import (
	"log"
	"net/http"

	"pokedex/helpers"
	"pokedex/services"
)

func List(w http.ResponseWriter, r *http.Request) {
	liste, err := services.GetAllPokemon()
	if err != nil {
		log.Println(err)
		http.Error(w, "Service momentanément indisponible.", http.StatusServiceUnavailable)
		return
	}
	helpers.Render(w, "list.html", liste)
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/pokemon", http.StatusMovedPermanently)
}
