package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(PokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + PokemonName

	if val, ok := c.cache.Get(url); ok {
		PokemonResp := Pokemon{}
		err := json.Unmarshal(val, &PokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return PokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	PokemonResp := Pokemon{}
	err = json.Unmarshal(data, &PokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return PokemonResp, nil
}
