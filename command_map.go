package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locResponse, err := cfg.apiClient.LocationList(cfg.nextLocURL)
	if err != nil {
		return err
	}

	cfg.nextLocURL = locResponse.Next
	cfg.prevLocURL = locResponse.Previous

	for _, loc := range locResponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocURL == nil {
		return errors.New("you're on the first page")
	}

	locResponse, err := cfg.apiClient.LocationList(cfg.prevLocURL)
	if err != nil {
		return err
	}

	cfg.nextLocURL = locResponse.Next
	cfg.prevLocURL = locResponse.Previous

	for _, loc := range locResponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
