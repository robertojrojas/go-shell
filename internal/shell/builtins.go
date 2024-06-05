package shell

import "github.com/sanurb/go-shell/cmd"

func exit(args []string) error {
	return cmd.Exit(args)
}

func echo(args []string) error {
	return cmd.Echo(args)
}

func pwd(args []string) error {
	return cmd.Pwd(args)
}

func cd(args []string) error {
	return cmd.Cd(args)
}

func cmdType(args []string) error {
	return cmd.Type(args)
}

func shellBuiltins() map[string]func(args []string) error {
	return map[string]func(args []string) error{
		"exit": exit,
		"echo": echo,
		"pwd":  pwd,
		"cd":   cd,
		"type": cmdType,
	}
}
