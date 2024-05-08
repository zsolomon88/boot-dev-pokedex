package main

import (
	"time"

	"github.com/zsolomon88/boot-dev-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 10 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	commandLineStart(cfg)
}
