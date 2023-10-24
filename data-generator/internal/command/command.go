package command

import (
	"github.com/arduino/go-paths-helper"
	"github.com/per1234/inoplatforms/data-generator/internal/registry"
	"github.com/per1234/inoplatforms/data-generator/internal/site"
	"github.com/spf13/cobra"
)

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

	dataObject := registry.Load(*paths.New(registryPath))

	dataObject.Populate()

	site.WritePages(dataObject, *paths.New(siteContentFolderPath))

	// TODO: Write search index data.

	site.WriteGeneratorData(dataObject, *paths.New(generatorDataPath))

	dataObject.Write(*paths.New(dataFilePath))
}
