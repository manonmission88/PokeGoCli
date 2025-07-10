# PokeGoCLI

A command-line Pokédex built in Go, designed to help users explore Pokémon locations, catch Pokémon, and inspect their stats directly from their terminal.

## Project Goal

The goal of this project is to create a lightweight and interactive CLI application that connects to the [PokeAPI](https://pokeapi.co/) to simulate a basic Pokémon-catching experience. The CLI allows you to explore the Pokémon world step-by-step, mimicking a minimal game loop — all in your terminal.

## Technologies Used

- **Go (Golang):** Main language used for building the CLI and handling HTTP requests, JSON unmarshalling, and program control.
- **PokeAPI:** Public API used to fetch data about Pokémon, locations, and stats.
- **Go Modules:** For dependency management.
- **Go Concurrency & Caching:** Implements basic in-memory caching to reduce repeated API calls.

## Features

- `map`: View a paginated list of available Pokémon locations.
- `mapb`: Navigate back to the previous page of locations.
- `explore <location>`: List Pokémon available in a given area.
- `catch <pokemon>`: Try to catch a Pokémon. Each Pokémon has a catch difficulty based on its base experience.
- `inspect <pokemon>`: View detailed stats and type(s) of any caught Pokémon.
- `help`: List all available commands.
- `exit`: Exit the Pokédex.


## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/PokeGoCli.git
    cd PokeGoCli
    ```

2. Run the application:

    ```bash
    go run .
    ```
## Example Usage

```bash
Pokedex > map
route-1
viridian-forest
Pokedex > map
route-1
viridian-forest
...
Pokedex > mapb
route-1
viridian-forest

Pokedex > explore viridian-forest
Exploring the viridian-forest....
Found Pokemon 
* pikachu
* caterpie
...

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats: 
  -speed: 90
  -attack: 55
  -defense: 40
Types: 
  -electric

...
Pokedex > pokedex
All Caught Pokemon:
- pikachu 

## Additional Resources

- [Boot.dev Go Track](https://boot.dev/learn/learn-golang): Interactive courses for learning Go fundamentals.
- [Go Official Documentation](https://golang.org/doc/): Comprehensive documentation for the Go programming language.
- [JSON-to-Go](https://mholt.github.io/json-to-go/): Convert JSON data to Go structs easily.

