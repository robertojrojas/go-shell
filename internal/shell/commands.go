package shell

import (
	"fmt"
)

func registerBuiltins() map[string]func(args []string) error {
	builtins := make(map[string]func(args []string) error)
	builtins["exit"] = exit
	builtins["echo"] = echo
	builtins["pwd"] = pwd
	builtins["cd"] = cd
	builtins["ls"] = ls
	builtins["type"] = cmdType
	return builtins
}

func cmdType(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no command provided")
	}
	fmt.Printf("%s is a builtin command\n", args[0])
	return nil
}
