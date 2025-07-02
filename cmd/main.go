package main

import (
	"os"

	"github.com/robertojrojas/go-shell/pkg/shell"
)

func main() {
	sh := shell.NewShell(os.Stdout)
	/*
		shellScript := `cat /etc/os-release
		echo "mini shell script works!"
		noop`
		cmdReader := strings.NewReader(shellScript)
		sh.Run(cmdReader, false)
	*/
	sh.Run(os.Stdin, true)
}
