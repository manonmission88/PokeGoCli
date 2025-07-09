package pokecache

import "time"

// field with creating time and value at that time
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
