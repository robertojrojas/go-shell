package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shell struct {
	builtins map[string]func(args []string) error
}

func New() *Shell {
	sh := &Shell{
		builtins: registerBuiltins(),
	}
	return sh
}

func (sh *Shell) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		sh.handleInput(strings.TrimSpace(input))
	}
}

func (sh *Shell) handleInput(input string) {
	if len(input) == 0 {
		return
	}

	args := strings.Split(input, " ")
	cmd := args[0]

	if function, found := sh.builtins[cmd]; found {
		sh.executeBuiltin(function, args[1:])
	} else {
		if err := sh.executeExternalCommand(cmd, args[1:]); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func (sh *Shell) executeBuiltin(function func(args []string) error, args []string) {
	if err := function(args); err != nil {
		fmt.Println("Error:", err)
	}
}
