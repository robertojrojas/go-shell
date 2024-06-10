package builtins

import (
	"github.com/sanurb/go-shell/internal/shell/commands"
	"os"
)

type ExitCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *ExitCommand) Execute() error {
	os.Exit(0)
	return nil
}
