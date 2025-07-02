package builtins

import (
	"fmt"
	"os"
	"strings"

	"github.com/robertojrojas/go-shell/pkg/shell/commands"
)

type CatCommand struct {
	commands.BaseCommand
	Args []string
}

func (c *CatCommand) Execute() error {
	fmt.Fprintln(c.Stdout, strings.Join(c.Args, " "))
	fdata, err := os.ReadFile(c.Args[0])
	if err != nil {
		return err
	}
	fmt.Fprintf(c.Stdout, "%s", string(fdata))
	return nil
}
