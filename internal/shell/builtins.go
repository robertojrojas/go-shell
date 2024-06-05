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

func ls(args []string) error { return cmd.Ls(args) }
