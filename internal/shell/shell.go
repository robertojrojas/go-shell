package shell

import (
	"bufio"
	"fmt"
	"github.com/sanurb/go-shell/internal/shell/parser"
	"github.com/sanurb/go-shell/internal/shell/registry"
	"os"
	"strings"
)

type Shell struct {
	parser *parser.CommandParser
}

func NewShell() *Shell {
	factory := registry.NewCommandFactory(os.Stdout)
	registry.RegisterBuiltins(factory)
	parser := parser.NewCommandParser(factory)
	return &Shell{
		parser: parser,
	}
}

func (sh *Shell) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		cmd, err := sh.parser.Parse(input)
		if err != nil {
			fmt.Println("Error parsing command:", err)
			continue
		}

		if cmd == nil {
			continue
		}

		if err := cmd.Execute(); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
