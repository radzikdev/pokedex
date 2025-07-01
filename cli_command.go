package main

import "fmt"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCliCommands () map[string]cliCommand {
return map[string]cliCommand {
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
}
}

func listCommands() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for k, v:= range getCliCommands() {
    fmt.Printf("%v: %v\n",k, v.description)
	}
	return nil
}