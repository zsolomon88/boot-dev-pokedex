package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type PokeLocArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetPokeLocArea(area *string) (PokeLocArea, error) {

	url := baseUrl + "/location-area/" + *area

	if val, ok := c.cache.Get(url); ok {
		locs := PokeLocArea{}
		jsonErr := json.Unmarshal(val, &locs)
		if jsonErr != nil {
			return PokeLocArea{}, jsonErr
		}
		fmt.Println("Cached result: ")
		return locs, nil
	}

	res, err := c.httpClient.Get(url)

	if err != nil {
		return PokeLocArea{}, err
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()
	if res.StatusCode > 299 {
		return PokeLocArea{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return PokeLocArea{}, err
	}

	locs := PokeLocArea{}
	jsonErr := json.Unmarshal(body, &locs)
	if jsonErr != nil {
		return PokeLocArea{}, err
	}

	c.cache.Add(url, body)
	return locs, nil
}