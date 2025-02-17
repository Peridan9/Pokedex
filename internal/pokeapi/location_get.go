package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation retrieves details of a specific location from the PokeAPI.
// If cached data is available, it is used instead of making an API request.
// Debug logs are included to track cache usage and potential parsing errors.
func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	// Check if the location data is available in the cache.
	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp) // Deserialize cached data.
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	// Create a new HTTP request to fetch location details.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	// Send the HTTP request using the client's HTTP client.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close() // Ensure response body is closed after reading.

	// Read response data.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	// Parse the JSON response into a Location struct.
	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}

	// Store the response data in the cache for future requests.
	c.cache.Add(url, data)
	return locationResp, nil
}
