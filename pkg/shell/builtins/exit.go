package builtins

import (
	"os"

	"github.com/robertojrojas/go-shell/pkg/shell/commands"
)

type ExitCommand struct {
	commands.BaseCommand
	Args []string
}

type SilentExitError struct{}

func (e *SilentExitError) Error() string {
	return ""
}

func (c *ExitCommand) Execute() error {
	os.Exit(0)
	return nil
}
