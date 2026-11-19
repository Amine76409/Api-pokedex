package main

import (
	"fmt"
	"pokedex/api"
)

func main() {
	pikachu, err := api.GetPokemon(25)
	fmt.Println(pikachu.Name.Fr, err)

	_, err = api.GetPokemon(9999)
	fmt.Println(err)

	liste, err := api.GetAllPokemon()
	fmt.Println(len(liste), err)
}
