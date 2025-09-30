package main

import "fmt"

func commandExplore(cfg *config, param string) error {
	if param == "" {
		return fmt.Errorf("error: no location area provided")
	}

	fmt.Printf("Exploring %s...\n", param)

	exploreResp, err := cfg.apiClient.ExploreList(param)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	encountersList := exploreResp.PokemonEncounters
	for _, pokemon := range encountersList {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
