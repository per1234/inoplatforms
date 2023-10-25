// Package root implements the data-generator root command.
package root

import (
	"github.com/arduino/go-paths-helper"
	"github.com/per1234/inoplatforms/data-generator/internal/data"
	"github.com/per1234/inoplatforms/data-generator/internal/site"
	"github.com/per1234/inoplatforms/registry/assets/go-registry/registry"
	"github.com/spf13/cobra"
)

// DataGenerator is the `data-generator` root command.
func DataGenerator(rootCommand *cobra.Command, cliArguments []string) {
	registryPath, err := rootCommand.Flags().GetString("registry")
	if err != nil {
		panic(err)
	}
	siteContentFolderPath, err := rootCommand.Flags().GetString("site-content")
	if err != nil {
		panic(err)
	}
	generatorDataPath, err := rootCommand.Flags().GetString("generator-data")
	if err != nil {
		panic(err)
	}
	dataFilePath, err := rootCommand.Flags().GetString("data-file")
	if err != nil {
		panic(err)
	}

	registryData := registry.Load(*paths.New(registryPath))

	dataData := data.Get(registryData)

	site.WritePages(dataData, *paths.New(siteContentFolderPath))

	// TODO: Write search index data.

	site.WriteGeneratorData(dataData, *paths.New(generatorDataPath))

	dataData.Write(*paths.New(dataFilePath))
}
