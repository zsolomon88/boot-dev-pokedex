package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}
type Cache struct {
	entry map[string]cacheEntry
	mux *sync.Mutex
	ttl time.Duration
}

func NewCache(ttl time.Duration) Cache {
	nCache := Cache{
		entry: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
		ttl: ttl,
	}
	go nCache.reapLoop()
	return nCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.entry[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if entry, ok := c.entry[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)

	for range ticker.C {
		fmt.Println("Cleaning up cache")
		c.mux.Lock()
		for key := range c.entry {
			if time.Since(c.entry[key].createdAt) > c.ttl {
				delete(c.entry, key)
			}
		}
		c.mux.Unlock()
	}
}