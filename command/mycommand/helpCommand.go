package mycommand

import (
	"fmt"

	"com.nokia.zayar/command"
)

type HelpCommand struct {
	command.BaseCommand
}

func (c *HelpCommand) Run(args []string) {
	fmt.Println("Available commands and their switches:")
	cmds := command.GetCommandList()
	for _, cmd := range cmds {
		fmt.Println(cmd.Usage())
	}
	fmt.Println("Enter a command or 'end' to quit:")
}

func (c *HelpCommand) Usage() string {
	return `
help or h: Show detailed explanations of commands and switches
`
}