package command

import (
	"fmt"
	"strings"
)

type Command interface {
	Name() string
	Description() string
	Run(args []string)
	Usage() string
}

type BaseCommand struct {
	name        string
	description string
}

func (bc *BaseCommand) Name() string {
	return bc.name
}

func (bc *BaseCommand) Description() string {
	return bc.description
}

func GetBaseCommand(name, description string) BaseCommand {
	return BaseCommand{name, description}
}

func PreprocessArgs(args []string) ([]string, error) {
	var processedArgs []string
	var currentArg string
	inQuotes := false

	for _, arg := range args {
		if strings.HasPrefix(arg, "\"") && strings.HasSuffix(arg, "\"") && len(arg) > 1 {
			processedArgs = append(processedArgs, strings.Trim(arg, "\""))
		} else if strings.HasPrefix(arg, "\"") {
			currentArg = strings.TrimPrefix(arg, "\"")
			inQuotes = true
		} else if strings.HasSuffix(arg, "\"") {
			if inQuotes {
				currentArg += " " + strings.TrimSuffix(arg, "\"")
				processedArgs = append(processedArgs, currentArg)
				currentArg = ""
				inQuotes = false
			} else {
				return nil, fmt.Errorf("mismatched quotes in argument: %s", arg)
			}
		} else if inQuotes {
			currentArg += " " + arg
		} else {
			processedArgs = append(processedArgs, arg)
		}
	}

	if inQuotes {
		return nil, fmt.Errorf("mismatched quotes detected")
	}

	return processedArgs, nil
}