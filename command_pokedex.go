package main

import (
	"fmt"
)

// commandPokedex prints the list of all caught Pokémon.
func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	// Iterate over caught Pokémon and print their names.
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
