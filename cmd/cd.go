package cmd

import (
	"fmt"
	"os"
)

func Cd(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cd: missing argument")
	}

	path, err := expandPath(args[0])
	if err != nil {
		return err
	}

	if err := os.Chdir(path); err != nil {
		return printError(args[0], err)
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
