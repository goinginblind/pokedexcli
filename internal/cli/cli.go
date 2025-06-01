package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/goinginblind/pokedexcli/internal/pokecache"
)

var cache *pokecache.Cache

// Runs the infinite loop until the user uses the 'quit' command as standard input
func Run() {
	cache = pokecache.NewCache(time.Duration(5 * time.Minute))
	cfg := &Config{Cache: cache}
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scan.Scan()
		gotString := scan.Text()

		if gotString == "" {
			continue
		}

		strSlice := cleanInput(gotString)
		prompt := strSlice[0]
		if command, ok := commands[prompt]; ok {
			if err := command.callback(cfg); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// Local function used to clean the standard user input
func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

var commands = map[string]CliCommand{}

// Initializes 'commands' map, used to not get int an infinite loop (since the help command itself needs the 'commands' map)
func init() {
	commands["help"] = CliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["exit"] = CliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["map"] = CliCommand{
		name:        "map",
		description: "Get next locations",
		callback:    printNextMapPage,
	}
	commands["mapb"] = CliCommand{
		name:        "mapb",
		description: "Get previous locations",
		callback:    printPrevMapPage,
	}
}
