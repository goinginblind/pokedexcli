package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	conf := config{}
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
			err := command.callback(&conf)
			fmt.Println(err)
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

// their functions
func commandExit(inp *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(inp *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// map functionality
type config struct {
	Next     string
	Previous string
}

func printNextMapPage(inp *config) error {
	var fetchUrl string
	if inp.Next == "" {
		fetchUrl = "https://pokeapi.co/api/v2/location-area/"
	} else {
		fetchUrl = inp.Next
	}
	res, err := http.Get(fetchUrl)
	if err != nil {
		return fmt.Errorf("could not get response body: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}

	var decodedRes LocationAreaResponse
	err = json.NewDecoder(res.Body).Decode(&decodedRes)
	if err != nil {
		return fmt.Errorf("could not decode response: %v", err)
	}

	for _, entry := range decodedRes.Results {
		fmt.Println(entry.Name)
	}

	inp.Next = decodedRes.Next
	inp.Previous = decodedRes.Previous

	return nil
}

func printPrevMapPage(inp *config) error {
	var fetchUrl string
	if inp.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		fetchUrl = inp.Previous
	}

	res, err := http.Get(fetchUrl)
	if err != nil {
		return fmt.Errorf("could not get response body: %v", err)
	}
	defer res.Body.Close()

	var decodedRes LocationAreaResponse
	err = json.NewDecoder(res.Body).Decode(&decodedRes)
	if err != nil {
		return fmt.Errorf("could not decode response: %v", err)
	}

	for _, entry := range decodedRes.Results {
		fmt.Println(entry.Name)
	}

	inp.Next = decodedRes.Next
	inp.Previous = decodedRes.Previous

	return nil
}
