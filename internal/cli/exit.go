package cli

import (
	"fmt"
	"os"
)

// Quits the inf loop
func commandExit(inp *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
