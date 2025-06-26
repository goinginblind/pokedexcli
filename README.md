# Pokedex CLI

A command-line interface (CLI) application written in Go for exploring Pokémon data using the PokéAPI.

## Features
- Explore and search for Pokémon
- Inspect Pokémon details
- Catch and manage your own Pokédex
- Map and explore different locations
- Caching for faster repeated queries

## Getting Started

### Prerequisites
- Go 1.18 or newer

### Installation
Clone the repository:
```bash
git clone github.com/goinginblind/pokedexcli
cd pokedexcli
```

Build the application:
```bash
go build -o pokedex
```

### Usage
Run the CLI:
```bash
./pokedex
```

Type `help` in the CLI for a list of available commands.

## Project Structure
- `main.go`: Entry point for the CLI application
- `internal/cli/`: CLI command implementations
- `internal/pokeapi/`: PokéAPI client and types
- `internal/pokecache/`: Caching utilities

## License
MIT License
