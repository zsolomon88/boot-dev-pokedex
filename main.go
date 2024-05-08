package main

import (
	"time"

	"github.com/zsolomon88/boot-dev-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 60 * time.Second)
	cfg := &config{
		pokedex: make(map[string]pokeapi.PokeInfo),
		pokeapiClient: pokeClient,
	}
	commandLineStart(cfg)
}
