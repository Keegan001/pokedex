package pokeapi

import (
	"encoding/json"
	"net/http"
)

type MapAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (*MapAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	areas := &MapAreas{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(areas)
	if err != nil {
		return nil, err
	}

	return areas, nil
}
