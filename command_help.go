package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range commandList() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}