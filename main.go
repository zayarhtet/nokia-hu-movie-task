package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"com.nokia.zayar/command"
	"com.nokia.zayar/command/mycommand"
	"com.nokia.zayar/repository"
	"com.nokia.zayar/service"
)

func main() {

	repository.Init("data.json")
	defer repository.Close()
	service.Init()

	command.RegisterCommand(&mycommand.ListCommand{
		BaseCommand: command.GetBaseCommand("l", "List command"),
	})
	command.RegisterCommand(&mycommand.AddCommand{
		BaseCommand: command.GetBaseCommand("a", "Add command"),
	})
	command.RegisterCommand(&mycommand.DeleteCommand{
		BaseCommand: command.GetBaseCommand("d", "Delete command"),
	})
	command.RegisterCommand(&mycommand.HelpCommand{
		BaseCommand: command.GetBaseCommand("help", "Show detailed explanations")})

	cmds := command.GetCommandList()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nAvailable commands:")
		for _, cmd := range cmds {
			fmt.Printf("  %s: %s\n", cmd.Name(), cmd.Description())
		}
		fmt.Print("Enter a command or 'end' to quit: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v", err)
			continue
		}

		fmt.Println()

		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "end" {
			fmt.Println("Exiting the program.")
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 { continue }

		commandName := args[0]
		cmd, found := command.GetCommand(commandName)
		if !found {
			log.Printf("Unknown command: %s", commandName)
			continue
		}

		preprocessedArgs, err := command.PreprocessArgs(args[1:])
		if err != nil {
			log.Printf("Error preprocessing args for command %s: %v", commandName, err)
            continue
		}
		cmd.Run(preprocessedArgs)
	}
}