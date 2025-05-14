package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"github.com/MadridMAC/go_pokedex/internal/pokeapi"
)

func runRepl(cfg *config) {
	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		userInput.Scan()

		output := cleanInput(userInput.Text())
		if len(output) == 0 {
			continue
		}

		command := output[0]
		val, ok := commandList()[command]
		if ok {
			// val.callback()
			err := val.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
}

func commandList() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Obtain the following page of Pokemon locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Obtain the previous page of Pokemon locations",
			callback: commandMapb,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

