package main

import (
	"fmt"
	"os"
)

func commandMap() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
