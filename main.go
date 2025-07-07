package main

import (
	"fmt"
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
	fmt.Println("Hello, World!")
	check := cleanInput("Charmander Bulbasaur PIKACHU")
	fmt.Println(check)

}
