package main

import ( 
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	mp := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name: "map",
			description: "Displays next 20 locations",
			callback: 
		}
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: displayHelp,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleanSlice := cleanInput(input)
		val, err := mp[strings.ToLower(cleanSlice[0])] 
		if err == true {
			val.callback()
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

