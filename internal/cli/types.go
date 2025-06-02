package cli

import (
	"github.com/goinginblind/pokedexcli/internal/pokeapi"
	"github.com/goinginblind/pokedexcli/internal/pokecache"
)

// Config contains 'Next' and 'Previous' fields with URLs, and 'Cache' (for now)
type Config struct {
	Next     string
	Previous string
	Caught   map[string]pokeapi.Pokemon
	Cache    *pokecache.Cache
}

// Command containing a callback to execute the command prompted by the user
type CliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}
