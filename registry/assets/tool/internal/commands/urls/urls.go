// Package urls implements the `urls` command.
package urls

func Urls(urlsCommand *cobra.Command, cliArguments []string) {
	registryPath, err := duplicatesCommand.Flags().GetString("registry")
	if err != nil {
		panic(err)
	}
	registryData := registry.Load(*paths.New(registryPath))

	for _, packageProvider := range registry.PackageProviders {
		// TODO: check if URLs don't load
	}
}
