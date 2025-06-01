package cli

import (
	"github.com/goinginblind/pokedexcli/internal/pokecache"
)

// Config contains 'Next' and 'Previous' fields with URLs, and 'Cache' with map pages (for now)
type Config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
}

// Command containing a callback to execute the command prompted by the user
type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}
