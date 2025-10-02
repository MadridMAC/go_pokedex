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

	baseCatchRate := 10
	if pokemonData.BaseExperience > 200 && pokemonData.BaseExperience <= 400 {
		baseCatchRate = 25
	} else if pokemonData.BaseExperience > 400 {
		baseCatchRate = 40
	}

	if catchAttempt := rand.Intn(100); catchAttempt >= baseCatchRate {
		fmt.Printf("%s was caught!\n", param)

		// Two loops to get the necessary values from the Pokemon
		statList := map[string]int{}
		for _, v := range pokemonData.Stats {
			statName := v.Stat.Name
			statValue := v.BaseStat
			statList[statName] = statValue
		}
		typeList := []string{}
		for _, v := range pokemonData.Types {
			typeList = append(typeList, v.Type.Name)
		}

		// Add caught Pokemon and its details to Pokedex
		cfg.caughtPkmn[param] = Pokemon{
			name:   param,
			height: pokemonData.Height,
			weight: pokemonData.Weight,
			stats:  statList,
			types:  typeList,
		}

		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", param)
	}

	return nil
}
