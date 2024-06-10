package builtins

import (
	"fmt"
	"github.com/sanurb/go-shell/internal/shell/commands"
	"io/fs"
	"os"
	"strings"
	"time"
)

type LsCommand struct {
	commands.BaseCommand
	Options LsOptions
	DirPath string
}

type LsOptions struct {
	ShowAll    bool
	LongFormat bool
}

func (c *LsCommand) Execute() error {
	if c.DirPath == "" {
		c.DirPath = "."
	}
	return c.listDirectory()
}

func (c *LsCommand) SetArgs(args []string) {
	c.Options, c.DirPath = parseArgs(args)
}

func parseArgs(args []string) (LsOptions, string) {
	var options LsOptions
	var dirPath string

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if strings.Contains(arg, "a") {
				options.ShowAll = true
			}
			if strings.Contains(arg, "l") {
				options.LongFormat = true
			}
		} else {
			dirPath = arg
		}
	}

	return options, dirPath
}

func (c *LsCommand) listDirectory() error {
	entries, err := os.ReadDir(c.DirPath)
	if err != nil {
		return fmt.Errorf("cannot open directory %s: %w", c.DirPath, err)
	}

	for _, entry := range entries {
		if !c.Options.ShowAll && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		if c.Options.LongFormat {
			c.printLongFormat(entry)
		} else {
			fmt.Fprintln(c.Stdout, entry.Name())
		}
	}

	return nil
}

func (c *LsCommand) printLongFormat(entry fs.DirEntry) {
	info, err := entry.Info()
	if err != nil {
		fmt.Fprintf(c.Stdout, "error getting info for %s: %v\n", entry.Name(), err)
		return
	}
	modTime := info.ModTime().Format(time.RFC3339)
	fmt.Fprintf(c.Stdout, "%s %10d %s %s\n", info.Mode(), info.Size(), modTime, entry.Name())
}
