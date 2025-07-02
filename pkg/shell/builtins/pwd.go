package builtins

import (
	"fmt"
	"os"

	"github.com/robertojrojas/go-shell/pkg/shell/commands"
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
