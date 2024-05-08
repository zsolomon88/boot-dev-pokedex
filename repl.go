package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zsolomon88/boot-dev-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient	pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
}

func commandLineStart(cfg *config) {
	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("PokeDex> ")
		scanStatus := inputScanner.Scan()
		if !scanStatus {
			break
		}

		if scanStatus {
			words := cleanInput(inputScanner.Text())
			if len(words) == 0 {
				continue
			}
			command := words[0]
			commandList := getCommandsList()

			if _, ok := commandList[command]; !ok {
				fmt.Printf("Command %s not found\n", command)
				continue
			}
			cmdError := commandList[command].callback(cfg)
			if cmdError != nil {
				fmt.Println(cmdError.Error())
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
		"map": {
			name:        "map",
			description: "Obtain the next list of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Obtain the previous list of locations",
			callback:    commandMapB,
		},
	}
}