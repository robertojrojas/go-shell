package shell

import (
	"fmt"
	"os"
	"os/exec"
)

func (sh *Shell) executeExternalCommand(cmd string, args []string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	if err := command.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError
		}
		if _, ok := err.(*exec.Error); ok {
			fmt.Printf("%s: command not found\n", cmd)
			return nil
		}
		return err
	}
	return nil
}
