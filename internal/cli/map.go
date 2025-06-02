package cli

import (
	"fmt"

	"github.com/goinginblind/pokedexcli/internal/pokeapi"
)

// Prints a list of all locations from the next page
func printNextMapPage(cfg *Config, args []string) error {
	var url string
	if cfg.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = cfg.Next
	}

	var res pokeapi.LocationAreaResponse
	if err := fetchAndCacheJSON(url, &res, cfg.Cache, pokeapi.FetchLocRes); err != nil {
		return err
	}

	for _, entry := range res.Results {
		fmt.Println(entry.Name)
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous

	return nil
}

// Prints a list of all locations from the previous page
func printPrevMapPage(cfg *Config, args []string) error {
	var url string
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = cfg.Previous
	}

	var res pokeapi.LocationAreaResponse
	if err := fetchAndCacheJSON(url, &res, cfg.Cache, pokeapi.FetchLocRes); err != nil {
		return err
	}

	for _, entry := range res.Results {
		fmt.Println(entry.Name)
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous

	return nil
}
