package cli

import (
	"fmt"

	"github.com/goinginblind/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("nothing to explore, no arguments passed")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	var res pokeapi.EncounterResponse
	if err := fetchAndCacheJSON(url, &res, cfg.Cache, pokeapi.FetchEncounters); err != nil {
		return fmt.Errorf("could not fetch encounters for location %q: %v", args[0], err)
	}

	for _, pokemon := range res.PokemonEncounters {
		fmt.Println(pokemon.Name)
	}

	return nil
}
