package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"
)

type LsOptions struct {
	ShowAll    bool
	LongFormat bool
}

func Ls(args []string) error {
	options, dirPath := parseArgs(args)

	if dirPath == "" {
		dirPath = "."
	}

	return listDirectory(dirPath, options)
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

func listDirectory(dirPath string, options LsOptions) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("cannot open directory %s: %w", dirPath, err)
	}

	for _, entry := range entries {
		if !options.ShowAll && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		if options.LongFormat {
			printLongFormat(entry, dirPath)
		} else {
			fmt.Println(entry.Name())
		}
	}

	return nil
}

func printLongFormat(entry fs.DirEntry, dirPath string) {
	info, err := entry.Info()
	if err != nil {
		fmt.Printf("error getting info for %s: %v\n", entry.Name(), err)
		return
	}
	modTime := info.ModTime().Format(time.RFC3339)
	fmt.Printf("%s %10d %s %s\n", info.Mode(), info.Size(), modTime, entry.Name())
}
