package main

import (
	"fmt"
)

func displayHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Show next 20 locations")
	fmt.Println("mapb: Show previous 20 locations")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}