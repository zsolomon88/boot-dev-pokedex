package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommandsList() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
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

func commandExit() error {
	os.Exit(0)
	return nil
}

func main() {
	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("PokeDex> ")
		scanStatus := inputScanner.Scan()
		if !scanStatus {
			break
		}
		if scanStatus {
			command := inputScanner.Text()
			commandList := getCommandsList()

			if _, ok := commandList[command]; !ok {
				fmt.Printf("Command %s not found\n", command)
				continue
			}
			commandList[command].callback()
		}
	}
}
