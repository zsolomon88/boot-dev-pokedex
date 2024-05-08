package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type PokeLocs struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetPokeLocations(pageUrl *string) (PokeLocs, error) {

	url := baseUrl + "/location"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		locs := PokeLocs{}
		jsonErr := json.Unmarshal(val, &locs)
		if jsonErr != nil {
			return PokeLocs{}, jsonErr
		}
		fmt.Println("Cached result: ")
		return locs, nil
	}

	res, err := c.httpClient.Get(url)

	if err != nil {
		return PokeLocs{}, err
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()
	if res.StatusCode > 299 {
		return PokeLocs{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return PokeLocs{}, err
	}

	locs := PokeLocs{}
	jsonErr := json.Unmarshal(body, &locs)
	if jsonErr != nil {
		return PokeLocs{}, err
	}

	c.cache.Add(url, body)
	return locs, nil
}