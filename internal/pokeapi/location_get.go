package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	fmt.Println("Fetching URL:", url) // Debug: Log the URL being fetched

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("Using cached data")
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			fmt.Println("Cache data error:", err) // Debug: Check if cache is corrupted
			return Location{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	fmt.Println("Raw Response Body:", string(data))

	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		fmt.Println("JSON Parse Error:", err)
		return Location{}, err
	}

	c.cache.Add(url, data)
	return locationResp, nil
}
