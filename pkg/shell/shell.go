package shell

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/robertojrojas/go-shell/pkg/shell/builtins"
	"github.com/robertojrojas/go-shell/pkg/shell/parser"
	"github.com/robertojrojas/go-shell/pkg/shell/registry"
)

type Shell struct {
	parser *parser.CommandParser
}

func NewShell(out io.WriteCloser) *Shell {
	factory := registry.NewCommandFactory(out)
	registry.RegisterBuiltins(factory)
	parser := parser.NewCommandParser(factory)
	return &Shell{
		parser: parser,
	}
}

func (sh *Shell) Run(in io.Reader, prompt bool) {
	reader := bufio.NewReader(in)
	for {
		if prompt {
			fmt.Print("$ ")
		}

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

		if _, ok := cmd.(*builtins.NoopCommand); ok {
			return
		}
	}
}
