package main

import (
	"time"
	"github.com/MadridMAC/go_pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: client,
	}
	runRepl(cfg)
}