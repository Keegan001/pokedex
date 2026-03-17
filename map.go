package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type mapAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getAreasBack(cfg *config) error {
	if cfg.prevUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(cfg.prevUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	areas := mapAreas{}
	decoder := json.NewDecoder(res.Body)
    err = decoder.Decode(&areas)
    if err != nil {
        return err
    }
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	cfg.nextUrl = areas.Next
	if areas.Previous != nil {
		// This is the "Type Assertion"
		if val, ok := areas.Previous.(string); ok {
			cfg.prevUrl = val
		}
	} else {
		// If it's null, we are on the first page
		cfg.prevUrl = ""
	}
	return nil 
}

func getAreas(cfg *config) error {
	res, err := http.Get(cfg.nextUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	areas := mapAreas{}
	decoder := json.NewDecoder(res.Body)
    err = decoder.Decode(&areas)
    if err != nil {
        return err
    }
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	cfg.nextUrl = areas.Next
	if areas.Previous != nil {
		// This is the "Type Assertion"
		if val, ok := areas.Previous.(string); ok {
			cfg.prevUrl = val
		}
	} else {
		// If it's null, we are on the first page
		cfg.prevUrl = ""
	}
	return nil 
}