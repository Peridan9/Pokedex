package main

import (
	"errors"
	"fmt"
)

// commandInspect displays details of a specific caught Pokémon.
func commandInspect(cfg *config, args ...string) error {
	// Ensure the user provides exactly one Pokémon name.
	if len(args) != 1 {
		return errors.New("you must provide a single pokemon name")
	}

	name := args[0]

	// Check if the Pokémon has been caught.
	if pokemon, exists := cfg.caughtPokemon[name]; exists {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)

		// Print the Pokémon's stats.
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}

		// Print the Pokémon's types.
		fmt.Println("Types:")
		for _, typ := range pokemon.Types {
			fmt.Printf("  - %s\n", typ.Type.Name)
		}
		return nil
	}

	// Inform the user if the Pokémon is not in their collection.
	fmt.Println("you have not caught that pokemon")
	return nil

}
