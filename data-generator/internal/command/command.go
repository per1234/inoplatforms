package command

import (
	"github.com/arduino/go-paths-helper"
	"github.com/per1234/inoplatforms/data-generator/internal/registry"
	"github.com/per1234/inoplatforms/data-generator/internal/site"
	"github.com/spf13/cobra"
)

// TODO: get these via CLI flags.
var registryPath paths.Path = paths.New("../registry/registry.yml")
var siteContentFolderPath paths.Path = paths.New("../site/content")
var generatorDataPath paths.Path = paths.New("../site/data/generator-data.json")
var dataPath paths.Path = paths.New("../site/static/inoplatforms-data.json")

func dataGenerator(rootCommand *cobra.Command) {
	dataObject := registry.Load(rootCommand.Flags().GetString("registry"))

	dataObject.Populate()

	site.WritePages(dataObject, rootCommand.Flags().GetString("site-content"))

	site.WriteGeneratorData(dataObject, rootCommand.Flags().GetString("generator-data"))

	dataObject.Write(rootCommand.Flags().GetString("data-file"))
}
