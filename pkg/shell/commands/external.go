package commands

import (
	"os/exec"
)

type ExternalCommand struct {
	BaseCommand
	name string
	args []string
}

func NewExternalCommand(name string, args []string) *ExternalCommand {
	return &ExternalCommand{name: name, args: args}
}

func (c *ExternalCommand) Execute() error {
	cmd := exec.Command(c.name, c.args...)
	cmd.Stdin = c.Stdin
	cmd.Stdout = c.Stdout
	cmd.Stderr = c.Stdout

	return cmd.Run()
}
