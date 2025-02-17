package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations retrieves a paginated list of locations from the PokeAPI.
// If a `pageURL` is provided, it fetches that specific page; otherwise, it starts from the base location endpoint.
// It first checks the cache before making an API request.
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL // Use provided page URL if available.
	}

	// Check if the location data is already cached.
	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp) // Deserialize cached data.
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	// Create a new HTTP request to fetch location data.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Send the HTTP request using the client's HTTP client.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close() // Ensure response body is closed after reading.

	// Read response data.
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Parse the JSON response into a RespShallowLocations struct.
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
