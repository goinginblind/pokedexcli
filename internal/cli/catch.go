package cli

import (
	"fmt"
	"math/rand/v2"

	"github.com/goinginblind/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("nothing to catch, no arguments passed")
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]

	var res pokeapi.Pokemon
	if err := fetchAndCacheJSON(url, &res, cfg.Cache, pokeapi.FetchPokemon); err != nil {
		return fmt.Errorf("could not fetch pokemon %q: %v", args[0], err)
	}

	chance := rand.IntN(390)
	if chance/(res.BaseExperience+1) >= 1 {
		fmt.Printf("%v was caught!\n", res.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.Caught[res.Name] = res
	} else {
		fmt.Printf("%v escaped!\n", res.Name)
	}

	return nil
}
