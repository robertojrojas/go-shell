package builtins

import (
	"github.com/robertojrojas/go-shell/pkg/shell/commands"
)

type NoopCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *NoopCommand) Execute() error {
	return nil
}
