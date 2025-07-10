package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/manonmission88/PokeGoCli/internal/pokecache"
	"github.com/manonmission88/PokeGoCli/pokeapi"
)

// split the sentence into list of words
func cleanInput(text string) []string {
	updatedText := strings.Split(text, " ")
	finalResult := []string{}
	for _, word := range updatedText {
		word = strings.ToLower(strings.TrimSpace(word))
		if word != "" {
			finalResult = append(finalResult, word)
		}
	}
	return finalResult

}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args []string) error
}

func commandMapf(cfg *config, args []string) error {
	resp, err := cfg.PokeClient.CallLocation(cfg.NextLocation)
	if err != nil {
		return err
	}
	// supporting paginations
	cfg.NextLocation = resp.Next
	cfg.PreviousLocation = resp.Previous
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.PreviousLocation == nil {
		return errors.New("no any previous page--you are on the very first page")
	}
	resp, err := cfg.PokeClient.CallLocation(cfg.PreviousLocation)
	if err != nil {
		return err
	}
	// supporting paginations
	cfg.NextLocation = resp.Next
	cfg.PreviousLocation = resp.Previous
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// method to exit the repl
func commandExit(cfg *config, args []string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// method for help
func commandHelp(cfg *config, args []string) error {
	fmt.Println("Available commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// method to return all the explored locations
func commandExplore(cfg *config, args []string) error {
	if len(args) > 1 {
		return errors.New("too many parameters : usage explore <area name>")
	}
	if len(args) == 0 {
		return errors.New("usage explore <area name>")
	}
	area := args[0]
	fmt.Println("Exploring the " + area + "....")
	resp, err := cfg.PokeClient.ExploreLocation(area)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon ")
	for _, loc := range resp.PokemonEncounters {
		fmt.Println("* " + loc.Pokemon.Name)
	}
	return nil

}

// mapping commands and features
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "next all locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "all locations Pokemon located",
			callback:    commandExplore,
		},
	}
}

func main() {
	cfg := &config{
		PokeClient: pokeapi.NewClient(5*time.Second, pokecache.NewCache(7*time.Second)),
	}
	// read the cli input
	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("Welcome to the Pokedex!\n")
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		inputText := reader.Text()
		words := cleanInput(inputText)
		if len(words) == 0 {
			continue
		}
		command := words[0]
		cmd, ok := getCommands()[command]
		if ok {
			err := cmd.callback(cfg, words[1:])
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}
