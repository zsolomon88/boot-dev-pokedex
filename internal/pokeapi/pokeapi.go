package pokeapi

import (
	"net/http"
	"time"

	"github.com/zsolomon88/boot-dev-pokedex/internal/pokecache"
)

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheTtl time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheTtl),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}