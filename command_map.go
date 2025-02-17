package main

import (
	"errors"
	"fmt"
)

// commandMapf retrieves and prints the next set of locations from the PokeAPI.
func commandMapf(cfg *config, args ...string) error {
	// Fetch the next page of locations using the PokeAPI client.
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	// Update the URLs for navigating between location pages.
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// Print each location name from the response.
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// commandMapb retrieves and prints the previous set of locations from the PokeAPI.
func commandMapb(cfg *config, args ...string) error {
	// Check if there is a previous page; if not, return an error.
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	// Fetch the previous page of locations using the PokeAPI client.
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	// Update the URLs for navigating between location pages.
	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	// Print each location name from the response.
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
