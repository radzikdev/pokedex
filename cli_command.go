package main

import (
	"fmt"
	"github.com/radzikdev/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error     
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    listCommands,
		},
		"map": {
			name:        "map",
			description: "Displays location areas",
			callback:    pokeapi.MapNextLocationAreas,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous location areas",
			callback:    pokeapi.MapPreviousLocationAreas,
		},
	}
}

func listCommands(*pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for k, v := range getCliCommands() {
		fmt.Printf("%v: %v\n", k, v.description)
	}
	return nil
}
