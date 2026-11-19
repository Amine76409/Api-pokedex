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
