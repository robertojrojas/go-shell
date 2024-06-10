package builtins

import (
	"fmt"
	"github.com/sanurb/go-shell/internal/shell/commands"
	"strings"
)

type EchoCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *EchoCommand) Execute() error {
	fmt.Fprintln(c.Stdout, strings.Join(c.Args, " "))
	return nil
}
