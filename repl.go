package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
}

func startRepl() {
	// Create NewScanner to receive Stdin input
	input := bufio.NewScanner(os.Stdin)
	// REPL proper; Scan and receive input, then clean it
	// If no text is received, continue. Otherwise, print the first	cleaned "word" as the command
	for {
		fmt.Print("Pokedex > ")
		input.Scan()

		receivedText := cleanInput(input.Text())
		if len(receivedText) == 0 {
			continue
		}
		commandText := receivedText[0]

		validCommand, ok := getCommands()[commandText]
		if ok {
			err := validCommand.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: invalid command")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
