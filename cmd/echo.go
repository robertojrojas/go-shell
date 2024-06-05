package cmd

import (
	"fmt"
	"strings"
)

func Echo(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}
