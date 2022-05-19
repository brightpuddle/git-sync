package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
)

// Populated from CI/CD
var version string

// Args are CLI arguments.
type Args struct {
	Path string `arg:"-p,--path" help:"Path to search for git repos"`
}

// Description is the CLI description string.
func (Args) Description() string {
	return `Git sync extention`
}

// Version is the CLI version string, populated by goreleaser.
func (Args) Version() string {
	if version == "" {
		return "development build"
	}
	return fmt.Sprintf("version %s", version)
}

// getArgs parses the CLI args.
func getArgs() Args {
	args := Args{}
	arg.MustParse(&args)
	if args.Path == "" {
		home := os.Getenv("HOME")
		args.Path = filepath.Join(home, "src")
	}
	return args
}
