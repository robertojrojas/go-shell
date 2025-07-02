package builtins

import (
	"fmt"
	"strings"

	"github.com/robertojrojas/go-shell/pkg/shell/commands"
)

type EchoCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *EchoCommand) Execute() error {
	fmt.Fprintln(c.Stdout, strings.Join(c.Args, " "))
	return nil
}
