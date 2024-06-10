package builtins

import (
	"fmt"
	"github.com/sanurb/go-shell/internal/shell/commands"
	"os"
)

type PwdCommand struct {
	commands.BaseCommand
}

func (c *PwdCommand) Execute() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintln(c.Stdout, dir)
	return nil
}
