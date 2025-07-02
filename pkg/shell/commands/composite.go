package commands

import (
	"io"
)

type CompositeCommand struct {
	commands []Command
}

func (c *CompositeCommand) Add(cmd Command) {
	c.commands = append(c.commands, cmd)
}

func (c *CompositeCommand) Execute() error {
	for _, cmd := range c.commands {
		if err := cmd.Execute(); err != nil {
			return err
		}
	}
	return nil
}

func (c *CompositeCommand) SetStdin(r io.Reader) {
	for _, cmd := range c.commands {
		cmd.SetStdin(r)
	}
}

func (c *CompositeCommand) SetStdout(w io.WriteCloser) {
	for _, cmd := range c.commands {
		cmd.SetStdout(w)
	}
}

func (c *CompositeCommand) StdinPipe() (io.WriteCloser, error) {
	return nil, nil
}

func (c *CompositeCommand) StdoutPipe() (io.ReadCloser, error) {
	return nil, nil
}
