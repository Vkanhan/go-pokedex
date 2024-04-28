package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Invalid command")
			continue
		}
	}
}

type cliCommand struct {
	name string
	description string
	callback func(*config)error
}

func getCommands()map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name : "help",
			description: "Prints the help menu",
			callback: callbackHelp,
		},
		"map": {
			name : "map",
			description: "Get the next page of locations",
			callback: callbackMap,
		},
		// "mapb": {
		// 	name : "mapb",
		// 	description: "Get the previous page of locations",
		// 	callback: commandMapb,
		// },
		"exit": {
			name : "exit",
			description : "Turns off the pokedox",
			callback: callBackExit,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
