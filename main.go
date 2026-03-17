package main

import ( 
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	sharedConf := &config{
				nextUrl:"https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
				prevUrl:"",
			}
	mp := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			conf: &config{
				nextUrl:"",
				prevUrl: "",
			},
		},
		"map": {
			name: "map",
			description: "Displays next 20 locations",
			callback: getAreas,
			conf: sharedConf,
		},
		"mapb": {
			name: "mapb",
			description: "Displays previous 20 locations",
			callback: getAreasBack,
			conf: sharedConf,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: displayHelp,
			conf: &config{
				nextUrl:"",
				prevUrl: "",
			},
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
			val.callback(val.conf)
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

