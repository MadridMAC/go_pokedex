package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreList(locArea string) (RespAreaDetail, error) {
	// error handling when no locArea is given
	// if locArea == "" {
	// 	return RespAreaDetail{}, fmt.Errorf("error: no location area provided")
	// }

	// Get URL from pokeapi.go, and append the locArea
	url := baseURL + "/location-area/" + locArea

	// Check if URL exists in cache; if so, unmarshal the cached data
	value, ok := c.cache.Get(url)
	if ok {
		cachedResp := RespAreaDetail{}
		if err := json.Unmarshal(value, &cachedResp); err != nil {
			return RespAreaDetail{}, err
		}
		return cachedResp, nil
	}

	// Create GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespAreaDetail{}, err
	}

	// Send GET request through c.httpClient as Client
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreaDetail{}, err
	}
	defer resp.Body.Close()

	// Read response data from GET request
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespAreaDetail{}, err
	}

	// Unmarshal received JSON data into Go struct
	exploreResp := RespAreaDetail{}
	if err := json.Unmarshal(data, &exploreResp); err != nil {
		return RespAreaDetail{}, err
	}

	// Add response data to cache for future use
	c.cache.Add(url, data)

	return exploreResp, nil

}
