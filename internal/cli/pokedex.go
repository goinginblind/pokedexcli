package cli

import "fmt"

func commandPokedex(cfg *Config, args []string) error {
	if len(cfg.Caught) == 0 {
		fmt.Println("You haven't caught any yet!")
	}
	fmt.Println("Your pokedex:")
	for i := range cfg.Caught {
		fmt.Printf(" - %v\n", i)
	}
	return nil
}
