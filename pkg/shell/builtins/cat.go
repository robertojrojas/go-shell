package builtins

import (
	"fmt"
	"os"

	"github.com/robertojrojas/go-shell/pkg/shell/commands"
)

type CatCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *CatCommand) Execute() error {
	decider := c.Args[0]
	fname := c.Args[0]
	silentExitError := false
	if decider == "-e" {
		silentExitError = true
		fname = c.Args[1]
	}

	fdata, err := os.ReadFile(fname)
	if err != nil {
		return err
	}
	fmt.Fprintf(c.Stdout, "%s", string(fdata))
	if silentExitError {
		return &SilentExitError{}
	}
	return nil
}
