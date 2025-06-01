package cli

import (
	"encoding/json"
	"fmt"

	"github.com/goinginblind/pokedexcli/internal/pokeapi"
)

// Prints a list of all locations from the next page
func printNextMapPage(inp *Config) error {
	var url string
	if inp.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = inp.Next
	}

	res, err := fetchOrCache(url)
	if err != nil {
		return err
	}

	for _, entry := range res.Results {
		fmt.Println(entry.Name)
	}

	inp.Next = res.Next
	inp.Previous = res.Previous

	return nil
}

// Prints a list of all locations from the previous page
func printPrevMapPage(inp *Config) error {
	var url string
	if inp.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = inp.Previous
	}

	res, err := fetchOrCache(url)
	if err != nil {
		return err
	}

	for _, entry := range res.Results {
		fmt.Println(entry.Name)
	}

	inp.Next = res.Next
	inp.Previous = res.Previous

	return nil
}

// Fetches or caches the map pages
func fetchOrCache(url string) (pokeapi.LocationAreaResponse, error) {
	var res pokeapi.LocationAreaResponse
	// check cache
	if raw, ok := cache.Get(url); ok {
		err := json.Unmarshal(raw, &res)
		if err != nil {
			return res, fmt.Errorf("%v: could not decode a response from cache", err)
		}
		// or fetch it from the url and then cache it
	} else {
		fetched, err := pokeapi.FetchLocRes(url)
		if err != nil {
			return res, fmt.Errorf("%v: could not fetch a response", err)
		}

		raw, err := json.Marshal(fetched)
		if err != nil {
			return res, fmt.Errorf("%v: could not marshal cache entry", err)
		}
		cache.Add(url, raw)
		res = *fetched
	}
	return res, nil
}
