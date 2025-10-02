package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, param string) error {
	if param == "" {
		return fmt.Errorf("error: no pokemon provided")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", param)

	pokemonData, err := cfg.apiClient.PokemonData(param)
	if err != nil {
		return err
	}

	baseCatchRate := 5
	if pokemonData.BaseExperience > 200 && pokemonData.BaseExperience <= 400 {
		baseCatchRate = 10
	} else if pokemonData.BaseExperience > 400 {
		baseCatchRate = 15
	}

	if catchAttempt := rand.Intn(20); catchAttempt >= baseCatchRate {
		fmt.Printf("%s was caught!\n", param)
		cfg.caughtPkmn[param] = Pokemon{
			name: param,
		}
	} else {
		fmt.Printf("%s escaped!\n", param)
	}

	return nil
}
