package main

import (
	"fmt"
)

func commandPokedex(cfg *config, param string) error {
	if len(cfg.caughtPkmn) == 0 {
		return fmt.Errorf("you have not caught any Pokemon yet")
	}
	fmt.Println("Your Pokedex:")
	for _, pkmn := range cfg.caughtPkmn {
		fmt.Printf("- %s\n", pkmn.name)
	}
	return nil
}
