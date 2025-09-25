package pokeapi

import (
	"net/http"
	"time"

	"github.com/MadridMAC/go_pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func MakeClient(timeout, interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
