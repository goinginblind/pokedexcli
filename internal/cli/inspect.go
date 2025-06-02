package cli

import "fmt"

func commandInspect(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("nothing to inspect, no arguments passed")
	}

	pokemon, ok := cfg.Caught[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for i := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
	}
	fmt.Println("Types:")
	for i := range pokemon.Types {
		fmt.Println("  " + pokemon.Types[i].Type.Name)
	}

	return nil
}
