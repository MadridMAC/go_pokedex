package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(inputURL *string) (AreaLocationResponse, error) {
	// Create URL variable
	url := targetURL + "/location-area"
	if inputURL != nil {
		url = *inputURL
	}

	// Create GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaLocationResponse{}, err
	}
	
	// Send GET request to PokeAPI
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaLocationResponse{}, err
	}

	// Read obtained JSON data from PokeAPI
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaLocationResponse{}, err
	}

	// Unmarshal JSON data into Go struct
	output := AreaLocationResponse{}
	err = json.Unmarshal(data, &output)
	if err != nil {
		return AreaLocationResponse{}, err
	}

	// Return decoded data and nil
	return output, nil

}