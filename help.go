package main

import (
	"fmt"
)

func displayHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}