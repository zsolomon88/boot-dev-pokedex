package main

import "fmt"

func commandHelp(cfg *config) error {
	commands := getCommandsList()
	helpText := `
Welcome to the PokeDex!
Usage:

`

	for _, command := range commands {
		helpText = helpText + command.name + ": " + command.description + "\n"
	}

	fmt.Println(helpText)
	return nil
}