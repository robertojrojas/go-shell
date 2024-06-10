package registry

import (
	"fmt"
	"github.com/sanurb/go-shell/internal/shell/builtins"
	"github.com/sanurb/go-shell/internal/shell/commands"
	"io"
)

type CommandFactory struct {
	commands map[string]func(args []string, stdout io.Writer) commands.Command
	stdout   io.Writer
}

func NewCommandFactory(stdout io.Writer) *CommandFactory {
	return &CommandFactory{
		commands: make(map[string]func(args []string, stdout io.Writer) commands.Command),
		stdout:   stdout,
	}
}

func (f *CommandFactory) Register(name string, factory func(args []string, stdout io.Writer) commands.Command) {
	f.commands[name] = factory
}

func (f *CommandFactory) Create(name string, args []string) (commands.Command, error) {
	if factory, found := f.commands[name]; found {
		return factory(args, f.stdout), nil
	}
	return nil, fmt.Errorf("command not found")
}

func RegisterBuiltins(factory *CommandFactory) {
	factory.Register("exit", func(args []string, stdout io.Writer) commands.Command {
		return &builtins.ExitCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}}
	})
	factory.Register("echo", func(args []string, stdout io.Writer) commands.Command {
		return &builtins.EchoCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}, Args: args}
	})
	factory.Register("pwd", func(args []string, stdout io.Writer) commands.Command {
		return &builtins.PwdCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}}
	})
	factory.Register("cd", func(args []string, stdout io.Writer) commands.Command {
		return &builtins.CdCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}, Args: args}
	})
	factory.Register("ls", func(args []string, stdout io.Writer) commands.Command {
		lsCommand := &builtins.LsCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}}
		lsCommand.SetArgs(args)
		return lsCommand
	})

	factory.Register("type", func(args []string, stdout io.Writer) commands.Command {
		return &builtins.TypeCommand{BaseCommand: commands.BaseCommand{Stdout: stdout}, Args: args}
	})
}
