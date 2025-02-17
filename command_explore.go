package main

import (
	"errors"
	"fmt"
)

// commandExplore fetches and displays information about a specified location,
// including Pokémon that can be encountered there.
func commandExplore(cfg *config, args ...string) error {

	// Ensure the user provides exactly one location name.
	if len(args) != 1 {
		return errors.New("you must provide a single location name")
	}

	name := args[0]

	// Fetch location details from the PokeAPI.
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	// Display location and encountered Pokémon.
	fmt.Printf("Exploring %s... \n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil

}
