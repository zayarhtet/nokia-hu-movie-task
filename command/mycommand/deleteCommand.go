package mycommand

import (
	"log"
	"flag"

	"com.nokia.zayar/command"
	"com.nokia.zayar/service"
)

type DeleteCommand struct {
	command.BaseCommand
}

func (c *DeleteCommand) Run(args []string) {
	deleteCmd := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	deletePersonString := deleteCmd.String("p", "", "Delete person")

	err := deleteCmd.Parse(args)
	if err != nil {
		log.Printf("Error parsing delete command: %v", err)
		return
	}

	service.DeleteItem(*deletePersonString)
}

func (c *DeleteCommand) Usage() string {
	return `
d: Delete item
  -p "name"     person name with double quoted
`
}