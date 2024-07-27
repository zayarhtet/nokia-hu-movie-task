package mycommand

import (
	"flag"
	"log"
	"strings"

	"com.nokia.zayar/command"
	"com.nokia.zayar/service"
)

type ListCommand struct {
	command.BaseCommand
}

func (c *ListCommand) Run(args []string) {
	listCmd := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	verboseFlag := listCmd.Bool("v", false, "List movies with starrings")
	tFilter := listCmd.String("t", "", "Filter movies by title (regex)")
	dFilter := listCmd.String("d", "", "Filter movies by director (regex)")
	aFilter := listCmd.String("a", "", "Filter movies by actor (regex)")
	ascFlag := listCmd.Bool("la", false, "List movies in ascending order")
	descFlag := listCmd.Bool("ld", false, "List movies in descending order")

	err := listCmd.Parse(args)
	if err != nil {
		log.Printf("Error parsing list command: %v", err)
		return
	}

	if *ascFlag && *descFlag {
		log.Println("The 'l' command does not support both -la and -ld switches.")
        return
	}

	*tFilter = strings.Trim(*tFilter, `"`)
	*dFilter = strings.Trim(*dFilter, `"`)
	*aFilter = strings.Trim(*aFilter, `"`)

	service.ListItems(*verboseFlag, *tFilter, *dFilter, *aFilter, *ascFlag, *descFlag)
}

func (c *ListCommand) Usage() string {
	return `
l: List items
  -v            List with details
  -t "regex"    Filter items by type (regex)
  -d "regex"    Filter items by date (regex)
  -a "regex"    Filter items by author (regex)
  -la           List items in ascending order
  -ld           List items in descending order
`
}