package command

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var commands = map[string]Command{}

func RegisterCommand(cmd Command) {
	commands[cmd.Name()] = cmd
}

func GetCommandList() []Command {
	var list []Command
	for _, cmd := range commands {
		list = append(list, cmd)
	}
	return list
}

func GetCommand(name string) (Command, bool) {
	val, found := commands[name]

	return val, found
}

func GetInputString() (string, bool) {
	input, err := reader.ReadString('\n')
    if err!= nil {
		log.Printf("Error reading input: %v", err)
		return "", false
    }

    return strings.TrimSpace(input), true
}

func GetInputInteger() (int, bool) {
	
	input, _ := GetInputString()
    birthYear, err := strconv.Atoi(input)
    if err != nil {
		log.Printf("Error reading input: %v", err)
        return 0, false
    }

	return birthYear, true
}