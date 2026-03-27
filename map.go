package main

import (
	"fmt"

	"github.com/Keegan001/pokedex/internal/pokeapi"
)

func getAreasBack(cfg *config) error {
	if cfg.prevUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	areas, err := pokeapi.GetLocationAreas(cfg.prevUrl)
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
	areas, err := pokeapi.GetLocationAreas(cfg.nextUrl)
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
