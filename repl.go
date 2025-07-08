package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/radzikdev/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := pokeapi.Config{}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command, exists := getCliCommands()[words[0]]
		if exists {
			err := command.callback(&config)
			if err != nil {
				fmt.Printf("command %v returned: %v", command.name, err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func commandExit(*pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
