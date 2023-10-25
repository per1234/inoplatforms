// Package duplicates implements the `duplicates` command.
package duplicates

import (
	"github.com/arduino/go-paths-helper"
	"github.com/per1234/inoplatforms/registry/assets/go-registry/registry"
	"github.com/spf13/cobra"
)

// Duplicates is the `duplicates` command.
func Duplicates(duplicatesCommand *cobra.Command, cliArguments []string) {
	registryPath, err := duplicatesCommand.Flags().GetString("registry")
	if err != nil {
		panic(err)
	}
	registryData := registry.Load(*paths.New(registryPath))

	for _, packageProvider := range registry.PackageProviders {
		// TODO: check for duplicate package provider IDs
		// TODO: check for duplicate package index URLs
		// TODO: check for duplicate platforms:
		// - no package index
		// - same packager name
		// - same architecture
		// - same source data
	}

}
