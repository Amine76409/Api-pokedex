package main

import (
	"log"
	"net/http"

	"pokedex/helpers"
	"pokedex/routes"
)

func main() {
	helpers.LoadTemplates()

	log.Println("Serveur lancé sur http://localhost:8080")
	err := http.ListenAndServe(":8080", routes.New())
	if err != nil {
		log.Fatal(err)
	}
}
