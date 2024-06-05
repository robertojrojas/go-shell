package cmd

import (
	"fmt"
	"os"
)

func Pwd(args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}
