package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		}, "help": {
			name:        "help",
			description: "Get Help on Pokedex",
			callback:    commandHelp,
		},
	}
	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		if !userInput.Scan() {
			break
		}
		inputLine := userInput.Text()
		input := cleanInput(inputLine)
		command, exists := commands[input[0]]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
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
	fmt.Println("  help: Displays a help message")
	fmt.Println("  exit: Exit the Pokedex")
	return nil
}

func cleanInput(text string) []string {
	words := []string{}
	word := ""
	for _, char := range text {
		if char == ' ' {
			if word != "" {
				words = append(words, strings.ToLower(word))
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
