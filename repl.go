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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args ...)
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

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}


type cliCommand struct {
	name string
	description string
	callback func(*config, ...string)error
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
			callback: commandMap,
		},
		"mapb": {
			name : "mapb",
			description: "Get the previous page of locations",
			callback: commandMapb,
		},
		"explore": {
			name : "explore{location area}",
			description: "List the pokemon in a location area",
			callback: commandExplore,
		},
		"catch": {
			name : "catch{pokemon name}",
			description: "Get the pokemon and add to pokedex",
			callback: commandCatch,
		},
		"inspect": {
			name : "inspect{pokemon name}",
			description: "Get the inforamtion about caught pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name : "pokedex{pokemon name}",
			description: "View all the pokemon in your pokedex",
			callback: commandPokedex,
		},
		"exit": {
			name : "exit",
			description : "Turns off the pokedox",
			callback: callBackExit,
		},
	}
}

