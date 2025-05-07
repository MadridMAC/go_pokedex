package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

func runRepl() {
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
			val.callback()
		} else {
			fmt.Println("Unknown command")
		}
		// fmt.Printf("Your command was: %s\n", command)
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
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

