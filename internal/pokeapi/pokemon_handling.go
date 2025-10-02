package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonData(pokemon string) (RespPokemonData, error) {
	// Get URL from pokeapi.go, and append the provided Pokemon
	url := baseURL + "/pokemon/" + pokemon

	// Check if URL exists in cache; if so, unmarshal the cached data
	value, ok := c.cache.Get(url)
	if ok {
		cachedResp := RespPokemonData{}
		if err := json.Unmarshal(value, &cachedResp); err != nil {
			return RespPokemonData{}, err
		}
		return cachedResp, nil
	}

	// Create GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonData{}, err
	}

	// Send GET request through c.httpClient as Client
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonData{}, err
	}
	defer resp.Body.Close()

	// Read response data from GET request
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonData{}, err
	}

	// Unmarshal received JSON data into Go struct
	pokemonResp := RespPokemonData{}
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return RespPokemonData{}, err
	}

	// Add response data to cache for future use
	c.cache.Add(url, data)

	return pokemonResp, nil

}
