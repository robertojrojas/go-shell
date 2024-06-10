package builtins

import (
	"fmt"
	"github.com/sanurb/go-shell/internal/shell/commands"
	"os"
)

type CdCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *CdCommand) Execute() error {
	if len(c.Args) == 0 {
		return fmt.Errorf("cd: missing argument")
	}

	path, err := expandPath(c.Args[0])
	if err != nil {
		return err
	}

	if err := os.Chdir(path); err != nil {
		return printError(c.Args[0], err)
	}

	return nil
}

func expandPath(path string) (string, error) {
	if path == "~" {
		return os.UserHomeDir()
	}
	return path, nil
}

func printError(arg string, err error) error {
	fmt.Printf("%s: No such file or directory\n", arg)
	return nil
}
