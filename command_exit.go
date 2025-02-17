package main

import (
	"fmt"
	"os"
)

// commandExit gracefully shuts down the application.
func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0) // Terminate the program.
	return nil // This line is unreachable but included for function signature consistency.
}
