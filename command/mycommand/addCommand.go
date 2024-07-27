package mycommand

import (
	"log"
	"flag"

	"com.nokia.zayar/command"
	"com.nokia.zayar/service"
)

type AddCommand struct {
	command.BaseCommand
}

func (c *AddCommand) Run(args []string) {
	addCmd := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	pFlag := addCmd.Bool("p", false, "Add person")
	mFlag := addCmd.Bool("m", false, "Add movie")

	err := addCmd.Parse(args)
	if err != nil {
		log.Printf("Error parsing add command: %v", err)
		return
	}

	if !*pFlag && !*mFlag {
		log.Println("The 'a' command requires at least one of -p or -m switches.")
		return
	}

	if *pFlag && *mFlag {
		log.Println("The 'a' command cannot use both -p and -m switches.")
		return
	}

	service.AddItem(*pFlag, *mFlag)
}

func (c *AddCommand) Usage() string {
	return `
a: Add item
  -p            adding for new person
  -m            adding for new movie
`
}