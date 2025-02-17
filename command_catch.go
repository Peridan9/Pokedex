package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// commandCatch attempts to catch a specified Pokémon.
// The catch success is determined randomly based on the Pokémon's base experience.
func commandCatch(cfg *config, args ...string) error {

	// Ensure the user provides exactly one Pokémon name.
	if len(args) != 1 {
		return errors.New("you must provide a single pokemon name")
	}

	name := args[0]

	// Check if the Pokémon has already been caught.
	if _, exists := cfg.caughtPokemon[name]; exists {
		fmt.Printf("You already caught %s!\n", name)
		return nil
	}

	// Fetch Pokémon details from the PokeAPI.
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// Generate a random value based on the Pokémon's base experience.
	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	// Determine if the Pokémon escapes or is caught.
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	// Successfully caught the Pokémon.
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	// Add the Pokémon to the user's caught Pokémon collection.
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil

}
