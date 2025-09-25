package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Add entries to the cache
func (c Cache) Add(key string, value []byte) {
	// Mutex lock for thread-safety
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// Get entries from the cache
func (c Cache) Get(key string) ([]byte, bool) {
	// Mutex lock for thread-safety
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	return entry.val, ok
}

// Reaploop method; uses a ticker to routinely clear the cache of older data
func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

// Reap method; loops through the cache and deletes entries considered old enough
func (c Cache) reap(curr time.Time, interval time.Duration) {
	// Mutex lock for thread-safety
	c.mu.Lock()
	defer c.mu.Unlock()
	// For each entry, if created past the provided interval, delete it
	for key, val := range c.cache {
		if val.createdAt.Before(curr.Add(-interval)) {
			delete(c.cache, key)
		}
	}

}

// Cache creation function; takes an interval for reaping purposes
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}
