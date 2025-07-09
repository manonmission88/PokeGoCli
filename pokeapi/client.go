package pokeapi

import (
	"net/http"
	"time"

	"github.com/manonmission88/PokeGoCli/internal/pokecache"
)

// custom client
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// creating new client constructor which takes the time duration and cache
func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}
