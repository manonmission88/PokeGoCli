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
	// first check on the cache
	if data, ok := c.cache.Get(url); ok {
		var localResp InnerLocations
		if err := json.Unmarshal(data, &localResp); err == nil {
			return localResp, nil
		}
	}
	// if cache hit
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
	c.cache.Add(url, data)
	locationResp := InnerLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return InnerLocations{}, err
	}
	return locationResp, nil

}

// pokemone - locations list
func (c *Client) ExploreLocation(innerLocation string) (AllInnerLocations, error) {
	url := baseUrl + "/location-area/" + innerLocation + "/"
	// first check on the cache
	if data, ok := c.cache.Get(url); ok {
		var localResp AllInnerLocations
		if err := json.Unmarshal(data, &localResp); err == nil {
			return localResp, nil
		}
	}
	// if cache hit
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AllInnerLocations{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AllInnerLocations{}, err
	}
	// unmarshall
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AllInnerLocations{}, err
	}
	c.cache.Add(url, data)
	locationResp := AllInnerLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return AllInnerLocations{}, err
	}
	return locationResp, nil

}

// catch the pokemon
// pokemone - returns the base experience
func (c *Client) CatchPokemon(pokemonName string) (PokemonStat, error) {
	url := baseUrl + "/pokemon/" + pokemonName + "/"
	// first check on the cache
	if data, ok := c.cache.Get(url); ok {
		var localResp PokemonStat
		if err := json.Unmarshal(data, &localResp); err == nil {
			return localResp, nil
		}
	}
	// if cache hit
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonStat{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonStat{}, err
	}
	// unmarshall
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonStat{}, err
	}
	c.cache.Add(url, data)
	locationResp := AllInnerLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return PokemonStat{}, err
	}
	return PokemonStat{}, nil

}
