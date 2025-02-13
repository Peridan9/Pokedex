package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	// Convert the text to lowercase
	text = strings.ToLower(text)
	// Remove the whitespace from start and end
	text = strings.TrimSpace(text)
	// Split the text into word based on whitespace between them and return a slice of all words
	return strings.Fields(text)
}
