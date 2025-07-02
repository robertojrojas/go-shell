package builtins

import (
	"os"

	"github.com/robertojrojas/go-shell/pkg/shell/commands"
)

type ExitCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *ExitCommand) Execute() error {
	os.Exit(0)
	return nil
}
