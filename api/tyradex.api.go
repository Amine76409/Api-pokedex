package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"pokedex/models"
)

var client = http.Client{Timeout: 5 * time.Second}

func GetAllPokemon() ([]models.Pokemon, error) {
	fmt.Println("Appel de la liste")

	req, err := http.NewRequest("GET", "https://tyradex.app/api/v1/pokemon", nil)
	if err != nil {
		return nil, models.ErrUnavailable
	}
	req.Header.Set("User-Agent", "pokedex-tp-ynov")

	rep, err := client.Do(req)
	if err != nil {
		return nil, models.ErrUnavailable
	}
	defer rep.Body.Close()

	fmt.Println("Code :", rep.StatusCode)
	if rep.StatusCode != 200 {
		return nil, models.ErrUnavailable
	}

	var liste []models.Pokemon
	json.NewDecoder(rep.Body).Decode(&liste)
	return liste, nil
}

func GetPokemon(id int) (models.Pokemon, error) {
	fmt.Println("Appel du pokémon", id)

	req, err := http.NewRequest("GET", "https://tyradex.app/api/v1/pokemon/"+strconv.Itoa(id), nil)
	if err != nil {
		return models.Pokemon{}, models.ErrUnavailable
	}
	req.Header.Set("User-Agent", "pokedex-tp-ynov")

	rep, err := client.Do(req)
	if err != nil {
		return models.Pokemon{}, models.ErrUnavailable
	}
	defer rep.Body.Close()

	fmt.Println("Code :", rep.StatusCode)
	if rep.StatusCode == 404 {
		return models.Pokemon{}, models.ErrNotFound
	}
	if rep.StatusCode != 200 {
		return models.Pokemon{}, models.ErrUnavailable
	}

	var pokemon models.Pokemon
	json.NewDecoder(rep.Body).Decode(&pokemon)
	return pokemon, nil
}
