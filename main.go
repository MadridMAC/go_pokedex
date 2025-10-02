package main

import (
	"time"

	"github.com/MadridMAC/go_pokedex/internal/pokeapi"
)

func main() {
	pokeapi_client := pokeapi.MakeClient(time.Second*5, time.Minute*5)
	cfg := &config{
		apiClient:  pokeapi_client,
		caughtPkmn: map[string]Pokemon{},
	}
	startRepl(cfg)
}
