package pokecache

import (
	"sync"
	"time"
)

// cacheEntry represents an individual cache record, storing the value and its creation timestamp.
type cacheEntry struct {
	createdAt time.Time // Timestamp when the cache entry was created
	val       []byte    // Cached data.
}

// Cache is a thread-safe in-memory cache that stores API responses temporarily.
type Cache struct {
	mux   *sync.Mutex           // Mutex to ensure thread-safe access.
	cache map[string]cacheEntry // Map to store cached responses by key.
}

// NewCache initializes a new Cache instance with a background cleanup loop.
// - `interval`: The time interval for clearing expired cache entries.
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry), // Initialize the cache map.
		mux:   &sync.Mutex{},               // Initialize the mutex.
	}

	// Start a background process to remove expired cache entries at regular intervals.
	go c.reapLoop(interval)

	return c
}

// Add stores a key-value pair in the cache with the current timestamp.
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(), // Store the time of caching.
		val:       value,
	}
}

// Get retrieves a value from the cache if it exists.
// Returns the cached value and a boolean indicating whether it was found.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok

}

// reapLoop runs a background cleanup process to remove expired cache entries at the specified interval.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

// reap removes cache entries that have exceeded the expiration duration.
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) { // Check if the entry is expired.
			delete(c.cache, k) // Remove expired entry from cache.
		}
	}
}
