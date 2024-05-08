package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zsolomon88/boot-dev-pokedex/internal/pokeapi"
)

type config struct {
	pokedex 		map[string]pokeapi.PokeInfo
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
			if commandList[command].args != len(words) {
				fmt.Printf("Invalid arguments, %s expects %d args\n", command, commandList[command].args)
				continue
			}
			cmdArgs := []string{}
			for i := 1; i < len(words); i++ {
				cmdArgs = append(cmdArgs, words[i])
			}
			cmdError := commandList[command].callback(cfg, cmdArgs...)
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
	callback    func(*config, ...string) error
	args		int
}

func getCommandsList() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			args:		 1,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			args:		 1,
		},
		"map": {
			name:        "map",
			description: "Obtain the next list of locations",
			callback:    commandMap,
			args:		 1,
		},
		"mapb": {
			name:        "mapb",
			description: "Obtain the previous list of locations",
			callback:    commandMapB,
			args:		 1,
		},
		"explore": {
			name:        "explore",
			description: "Explore a given location",
			callback:    commandExplore,
			args: 		 2,
		},
		"info": {
			name:        "info",
			description: "Obtain info about a Pokemon",
			callback:    commandInfo,
			args: 		 2,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
			args: 		 2,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon",
			callback:    commandInspect,
			args: 		 2,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List your Pokedex",
			callback:    commandPokedex,
			args: 		 1,
		},
	}
}