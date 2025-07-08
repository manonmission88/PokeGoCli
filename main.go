package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	callback    func() error
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
	}
}

// method to exit the repl
func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// method for help
func commandHelp() error {
	fmt.Println("Available commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func main() {
	// read the cli input
	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("Welcome to the Pokedex!\n")
	for {
		reader.Scan()
		inputText := reader.Text()
		words := cleanInput(inputText)
		if len(words) == 0 {
			continue
		}
		command := words[0]
		cmd, ok := getCommands()[command]
		if ok {
			err := cmd.callback()
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
