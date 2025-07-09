package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex // mutex to protect concurrent sessions
	ttl     time.Duration
}

// creating the new cache and starting the reap loop
func NewCache(interval time.Duration) *Cache {
	// add new entry
	c := &Cache{
		entries: make(map[string]cacheEntry),
		ttl:     interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) cacheEntry {
	// lock the cache
	c.mu.Lock()
	defer c.mu.Unlock()
	// adding key, val to the cache Entry
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, stat := c.entries[key]
	if !stat {
		return nil, false
	}
	// check the duration of created and range of time interval
	if time.Since(entry.createdAt) > c.ttl {
		delete(c.entries, key)
		return nil, false
	}
	return entry.val, true

}

// reap loop -> delete the entries greater than the given time duration
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entries {
			if time.Since(v.createdAt) > c.ttl {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}
