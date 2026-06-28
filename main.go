package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	var commands map[string]cliCommand

	commandHelp := func() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println()

		for _, cmd := range commands {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
		return nil
	}

	commands = map[string]cliCommand{
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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())

		if len(words) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}

		commandName := strings.ToLower(words[0])

		command, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			fmt.Print("Pokedex > ")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Print("Pokedex > ")
	}
}
