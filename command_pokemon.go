package main

import (
	"fmt"
	"math/rand"
)

func commandInfo(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("invalid arguments received")
	}

	fmt.Printf("Info for Pokemon: %s\n", args[0])
	pokemonInfo, infoErr := cfg.pokeapiClient.GetPokeInfo(&args[0])
	if infoErr != nil {
		return infoErr
	}
	fmt.Printf(" Name: %s\n", pokemonInfo.Name)
	fmt.Printf(" Base XP: %d\n", pokemonInfo.BaseExperience)
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("invalid arguments received")
	}

	pokemonInfo, infoErr := cfg.pokeapiClient.GetPokeInfo(&args[0])
	if infoErr != nil {
		return infoErr
	}

	fmt.Printf("Throwing a Pokeball at %s\n", pokemonInfo.Name)
	catchRate := rand.Intn(pokemonInfo.BaseExperience)

	if catchRate >= (pokemonInfo.BaseExperience / 4) {
		fmt.Printf("%s was caught!\n", pokemonInfo.Name)
		cfg.pokedex[pokemonInfo.Name] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
	}

	return nil
}

func commandInspect(cfg *config, args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("invalid arguments received")
	}

	if pokemon, ok := cfg.pokedex[args[0]]; !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf(`
Name: %s
Height: %d
Weight: %d
Stats:
`, pokemon.Name, pokemon.Height, pokemon.Weight)
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types: ")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokeType.Type.Name)
		}
	}	
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex: ")
	if len(cfg.pokedex) == 0 {
		fmt.Println("You have not caught any Pokemon")
		return nil
	}

	for key := range cfg.pokedex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}