package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
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
			err := command.callback()
			fmt.Println(err)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
