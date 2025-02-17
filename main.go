package main

import (
	// Import the internal PokeAPI package.
	// If cloning the repository, update the import path to match your local module structure.
	"time"

	"github.com/Peridan9/Pokedex/internal/pokeapi"
)

func main() {
	// Create a new PokeAPI client with a request timeout of 5 seconds
	// and a cache expiration time of 5 minutes.
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	// Initialize the configuration struct with an empty map for caught Pokémon
	// and the PokeAPI client for making API requests.
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{}, // Stores Pokémon caught by the user
		pokeapiClient: pokeClient,                   // API client for fetching Pokémon data
	}

	// Start the Read-Eval-Print Loop (REPL) to handle user input.
	startRepl(cfg)
}
