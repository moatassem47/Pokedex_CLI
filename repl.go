package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)


	for {
		fmt.Print("Pokedex>>")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		AvailableCommands := getCommands()
		command, ok := AvailableCommands[commandName]
		if ok == false {
			fmt.Printf("\"%v\" command is not found\n", commandName)
			continue
		}
		command.callback()
	
	}



}

type CliCommand struct {
	name        string
	description string
	callback    func()
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {name: "help", description: "view manual for the tool", callback: HelpCommand},
		"exit": {name: "exit", description: "kills the tool process", callback: ExitCommand},
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
