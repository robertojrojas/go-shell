package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

var builtins = map[string]bool{
	"exit": true,
	"echo": true,
	"pwd":  true,
	"cd":   true,
	"type": true,
}

func Type(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no command specified")
	}
	cmd := args[0]

	if _, found := builtins[cmd]; found {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return nil
	}

	path, err := exec.LookPath(cmd)
	if err == nil {
		absPath, _ := filepath.Abs(path)
		fmt.Printf("%s is %s\n", cmd, absPath)
		return nil
	}

	fmt.Printf("%s not found\n", cmd)
	return nil
}
