// Package cli defines the data-generator command line interface.
package cli

import (
	"github.com/spf13/cobra"
)

// Root creates a new data-generator command root.
func Root() *cobra.Command {
	rootCommand := &cobra.Command{
		Short:                 "Platforms data generator",
		Long:                  "data-generator generates data on Arduino Boards Platforms.",
		DisableFlagsInUseLine: false,
		Use:                   "data-generator",
		Run:                   command.DataGenerator,
	}

	rootCommand.PersistentFlags().String("registry", "", "Path to the registry file.")
	rootCommand.PersistentFlags().String("site-content", "", "Path to the website source content folder.")
	rootCommand.PersistentFlags().String("generator-data", "", "Path where the generator data file should be created.")
	rootCommand.PersistentFlags().String("data-file", "", "Path where the data file should be created.")

	return rootCommand
}
