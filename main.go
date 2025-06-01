package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/goinginblind/pokedexcli/internal/pokeapi"
	"github.com/goinginblind/pokedexcli/internal/pokecache"
)

var cache *pokecache.Cache

func main() {
	scan := bufio.NewScanner(os.Stdin)
	conf := config{}
	cache = pokecache.NewCache(100 * time.Second)
	for {
		fmt.Print("Pokedex > ")
		scan.Scan()
		gotString := scan.Text()

		if gotString == "" {
			continue
		}

		strSlice := CleanInput(gotString)
		prompt := strSlice[0]
		if command, ok := commands[prompt]; ok {
			if err := command.callback(&conf); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// Regular functions for main
func CleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

// Command and registry for them
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commands = map[string]cliCommand{}

// initializes commands map
func init() {
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Get next locations",
		callback:    printNextMapPage,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Get previous locations",
		callback:    printPrevMapPage,
	}
}

// quits
func commandExit(inp *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// prints a help message
func commandHelp(inp *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// config contains two fields with urls for prev and next pages
type config struct {
	Next     string
	Previous string
}

// print next page
func printNextMapPage(inp *config) error {
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

// print prev page
func printPrevMapPage(inp *config) error {
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
