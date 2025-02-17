package pokeapi

import (
	"net/http"
	"time"

	"github.com/Peridan9/Pokedex/internal/pokecache"
)

// Client represents an API client for interacting with the PokeAPI.
// It includes an HTTP client for making requests and a cache to store responses.
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient creates a new PokeAPI client with a specified request timeout and cache expiration interval.
// - `timeout`: Maximum duration for an API request before timing out.
// - `cacheInterval`: Duration after which cached responses are removed.
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval), // Initialize cache with expiration interval.
		httpClient: http.Client{
			Timeout: timeout, // Set the HTTP request timeout.
		},
	}
}
