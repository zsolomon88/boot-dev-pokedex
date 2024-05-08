package main

import (
	"fmt"
)


func commandMap(cfg *config) error {
	pokeLocsList, locsErr := cfg.pokeapiClient.GetPokeLocations(cfg.nextLocationUrl)
	if locsErr != nil {
		return locsErr
	}
	for _, pokeLoc := range pokeLocsList.Results {
		fmt.Println(pokeLoc.Name)
	}
	cfg.nextLocationUrl = &pokeLocsList.Next
	cfg.prevLocationUrl = &pokeLocsList.Previous
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationUrl == nil || *(cfg.prevLocationUrl) == "" {
		return fmt.Errorf("you're on the first page")
	}
	pokeLocsList, locsErr := cfg.pokeapiClient.GetPokeLocations(cfg.prevLocationUrl)
	if locsErr != nil {
		return locsErr
	}
	for _, pokeLoc := range pokeLocsList.Results {
		fmt.Println(pokeLoc.Name)
	}
	cfg.nextLocationUrl = &pokeLocsList.Next
	cfg.prevLocationUrl = &pokeLocsList.Previous
	return nil
}