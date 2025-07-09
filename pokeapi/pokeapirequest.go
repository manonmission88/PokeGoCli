package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	baseUrl = "https://pokeapi.co/api/v2"
)

// locations list
func (c *Client) CallLocation(pageUrl *string) (InnerLocations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return InnerLocations{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return InnerLocations{}, err
	}
	// unmarshall
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return InnerLocations{}, err
	}
	locationResp := InnerLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return InnerLocations{}, err
	}
	return locationResp, nil

}
