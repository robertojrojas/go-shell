package main

import (
	"os"
	"strings"

	"github.com/robertojrojas/go-shell/pkg/shell"
)

func main() {
	sh := shell.NewShell(os.Stdout)
	sh.Run(os.Stdin, true)

	runScript()
}

func runScript() {
	sh := shell.NewShell(os.Stdout)
	shellScript := `cat /etc/os-release
		echo "mini shell script works!"
		noop`
	cmdReader := strings.NewReader(shellScript)
	sh.Run(cmdReader, false)
}
