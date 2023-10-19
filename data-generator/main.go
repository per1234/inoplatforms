// Package main is the main package.
package main

import (
	"os"

	"log/slog"

	"github.com/per1234/inoplatforms/data-generator/internal/cli"
)

func main() {
	rootCommand := cli.Root()
	if err := rootCommand.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
