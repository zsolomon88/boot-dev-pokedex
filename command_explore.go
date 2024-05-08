package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("invalid arguments received")
	}

	fmt.Printf("Exploring: %s\n", args[0])
	pokeLocsList, locsErr := cfg.pokeapiClient.GetPokeLocArea(&args[0])
	if locsErr != nil {
		return locsErr
	}
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range pokeLocsList.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}