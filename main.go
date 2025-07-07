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
func main() {
	// read the cli input
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		inputText := reader.Text()
		words := cleanInput(inputText)
		if len(words) > 0 {
			fmt.Printf("Your command was: %s\n", words[0])
		} else {
			fmt.Println("Please type a command.")
		}
	}

}
