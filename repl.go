package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MadridMAC/go_pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
}

func startRepl(cfg *config) {
	// Create NewScanner to receive Stdin input
	input := bufio.NewScanner(os.Stdin)
	// REPL proper; Scan and receive input, then clean it. If no text is received, continue.
	// Otherwise, get the first two words as command and argument.
	for {
		fmt.Print("Pokedex > ")
		input.Scan()

		receivedText := cleanInput(input.Text())
		if len(receivedText) == 0 {
			continue
		}
		commandText := receivedText[0]
		parameter := ""
		if len(receivedText) > 1 {
			parameter = receivedText[1]
		}

		validCommand, ok := getCommands()[commandText]
		if ok {
			err := validCommand.callback(cfg, parameter)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: invalid command")
			continue
		}
	}
}

type config struct {
	apiClient  pokeapi.Client
	nextLocURL *string
	prevLocURL *string
	caughtPkmn map[string]Pokemon
}

// to be fleshed out
type Pokemon struct {
	name string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas, with each successive call displaying the next 20.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Takes the name of a location area as an argument. Returns a list of all Pokemon in the specified area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes the name of a Pokemon. Attempts to catch the Pokemon, adding it to the Pokedex if successful.",
			callback:    commandCatch,
		},
	}
}
