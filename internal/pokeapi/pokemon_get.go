package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemon retrieves details of a Pokémon by name from the PokeAPI.
// It first checks the cache for stored data. If not found, it makes an API request.
// The response is then cached for future use.
func (c *Client) GetPokemon(PokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + PokemonName

	// Check if the Pokémon data is available in the cache.
	if val, ok := c.cache.Get(url); ok {
		PokemonResp := Pokemon{}
		err := json.Unmarshal(val, &PokemonResp) // Deserialize cached data.
		if err != nil {
			return Pokemon{}, err
		}
		return PokemonResp, nil
	}

	// Create a new HTTP request to fetch Pokémon data.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	// Send the HTTP request using the client's HTTP client.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close() // Ensure response body is closed after reading.

	// Read response data.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Parse the JSON response into a Pokemon struct.
	PokemonResp := Pokemon{}
	err = json.Unmarshal(data, &PokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	// Store the response data in the cache for future requests.
	c.cache.Add(url, data)
	return PokemonResp, nil
}
