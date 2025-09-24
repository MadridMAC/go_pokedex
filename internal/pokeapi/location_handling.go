package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationList(givenURL *string) (RespLocationArea, error) {
	// Get URL from pokeapi.go | Modify with givenURL if provided
	url := baseURL + "/location-area"
	if givenURL != nil {
		url = *givenURL
	}

	// Create GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	// Send GET request through c.httpClient as Client
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	// Read response data from GET request
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	// Unmarshal received JSON data into Go struct
	locResponse := RespLocationArea{}
	if err := json.Unmarshal(data, &locResponse); err != nil {
		return RespLocationArea{}, err
	}

	return locResponse, nil

}
