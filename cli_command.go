package main

import "fmt"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
var cliCommands map[string]cliCommand
func initCliCommands () {
cliCommands = map[string]cliCommand {
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
	for k, v:= range cliCommands {
    fmt.Printf("%v: %v\n",k, v.description)
	}
	return nil
}