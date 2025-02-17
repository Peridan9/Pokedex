package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Peridan9/Pokedex/internal/pokeapi"
)

// cliCommand represents a command in the REPL (Read-Eval-Print Loop).
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error // Function to execute the command
}

// config holds global data that is required across different parts of the program.
type config struct {
	caughtPokemon    map[string]pokeapi.Pokemon // Stores caught Pokémon by name
	pokeapiClient    pokeapi.Client             // Client for making API requests
	nextLocationsURL *string                    // URL for fetching the next set of locations
	prevLocationsURL *string                    // URL for fetching the previous set of locations
}

// getCommands initializes and returns a map of available commands with their respective callbacks.
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all caught pokemons",
			callback:    commandPokedex,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:    commandInspect,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

// startRepl starts the Read-Eval-Print Loop (REPL) for user interaction.
func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin) // Scanner to read user input from standard input
	fmt.Println("Welcome to the Pokedex!")

	// Infinite loop until the user enters 'exit'
	for {
		fmt.Print("Pokedex > ") // Prompt for user input
		scanned := scanner.Scan()

		if !scanned { // If scanning fails (e.g., EOF), break the loop
			break
		}

		input := scanner.Text() // Read user input
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:] // Extract command arguments
		}

		// Check if the command exists in the available commands
		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println("Error:", err) // Print any command execution errors
			}
			continue
		} else {
			fmt.Println("Unknown command") // Handle unrecognized commands
			continue
		}

	}
}

// cleanInput processes user input by converting it to lowercase and splitting it into words.
func cleanInput(text string) []string {
	output := strings.ToLower(text) // Convert input to lowercase
	words := strings.Fields(output) // Split input into words based on spaces
	return words
}
