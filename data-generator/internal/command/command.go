package command

import (
	"github.com/per1234/inoplatforms/data-generator/internal/registry"
	"github.com/per1234/inoplatforms/data-generator/internal/site"
	"github.com/spf13/cobra"
)

func DataGenerator(rootCommand *cobra.Command) {
	dataObject := registry.Load(rootCommand.Flags().GetString("registry"))

	dataObject.Populate()

	site.WritePages(dataObject, rootCommand.Flags().GetString("site-content"))

	// TODO: Write search index data.

	site.WriteGeneratorData(dataObject, rootCommand.Flags().GetString("generator-data"))

	dataObject.Write(rootCommand.Flags().GetString("data-file"))
}
