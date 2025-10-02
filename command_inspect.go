package main

import "fmt"

func commandInspect(cfg *config, param string) error {
	pokemon, ok := cfg.caughtPkmn[param]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.name)
	fmt.Printf("Height: %v\n", pokemon.height)
	fmt.Printf("Weight: %v\n", pokemon.weight)
	fmt.Printf("Stats:\n")
	for key, value := range pokemon.stats {
		fmt.Printf("  - %s: %v\n", key, value)
	}
	fmt.Printf("Types:\n")
	for _, pkType := range pokemon.types {
		fmt.Printf("  - %v\n", pkType)
	}

	return nil
}
