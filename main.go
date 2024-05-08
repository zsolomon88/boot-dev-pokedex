package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
		"map": {
			name:        "map",
			description: "Obtain a list of locations",
			callback:    commandMap,
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
type PokeLocs struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var prevUrl string
var nextUrl string
func commandMap() error {
	url := "https://pokeapi.co/api/v2/location/"
	if nextUrl != nil {
		url = nextUrl
	}

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}
	
	locs := PokeLocs{}
	jsonErr := json.Unmarshal(body, &locs)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	prevUrl = locs.Previous
	nextUrl = locs.Next
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
