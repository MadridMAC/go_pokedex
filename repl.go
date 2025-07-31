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

		fmt.Printf("Your command was: %s\n", commandText)
	}
}
