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
		fmt.Printf("Your command was: %v\n", strSlice[0])
	}
}

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
