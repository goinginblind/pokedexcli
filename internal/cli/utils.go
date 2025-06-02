package cli

import (
	"encoding/json"
	"fmt"

	"github.com/goinginblind/pokedexcli/internal/pokecache"
)

// Fetches or caches (or both)
func fetchAndCacheJSON(url string, target any, cache *pokecache.Cache, fetchFunc func(string) (any, error)) error {
	if raw, ok := cache.Get(url); ok {
		return json.Unmarshal(raw, target)
	}

	fetched, err := fetchFunc(url)
	if err != nil {
		return fmt.Errorf("fetch failed: %w", err)
	}

	raw, err := json.Marshal(fetched)
	if err != nil {
		return fmt.Errorf("marshalling failed: %w", err)
	}

	cache.Add(url, raw)
	return json.Unmarshal(raw, target)
}
