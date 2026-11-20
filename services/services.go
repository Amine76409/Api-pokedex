package services

import (
	"pokedex/api"
	"pokedex/models"
)

func GetAllPokemon() ([]models.Pokemon, error) {
	liste, err := api.GetAllPokemon()
	if err != nil {
		return nil, err
	}

	var resultat []models.Pokemon
	for _, p := range liste {
		if p.ID != 0 {
			resultat = append(resultat, p)
		}
	}
	return resultat, nil
}
