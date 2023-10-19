package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/arduino/go-paths-helper"
	"gopkg.in/yaml.v3"
)

type RepositoryType struct {
	Url  string
	Ref  string
	Path string
}
type SourceType struct {
	Repository RepositoryType
}

type ToolDependencyType struct {
	Packager          string
	Name              string
	Version           string
	PackageProviderID string
}

type PlatformType struct {
	Name                   string
	Architecture           string
	Help                   string
	Boards                 []BoardType
	Source                 SourceType
	InstallationReferences []string             `yaml:"installation-references"`
	ToolsDependencies      []ToolDependencyType `yaml:"tools-dependencies"`
	DiscoveryDependencies  []ToolType
	MonitorDependencies    []ToolType
	Releases               []PlatformReleaseType
}

type BoardType struct {
	Name string
}

type PlatformReleaseType struct {
	Name                  string
	Version               string
	Help                  string
	ArchiveURL            string
	Checksum              string
	Boards                []BoardType
	ToolsDependencies     []ToolDependencyType `yaml:"tools-dependencies"`
	DiscoveryDependencies []ToolType
	MonitorDependencies   []ToolType
}

type SystemType struct {
	Checksum   string
	Host       string
	ArchiveURL string
}

type ToolType struct {
	Name     string
	Source   SourceType
	Releases []ToolReleaseType
}

type ToolReleaseType struct {
	Version string
	Systems []SystemType
}

type PackageType struct {
	Name       string
	Maintainer string
	WebsiteURL string
	Platforms  []PlatformType
	Tools      []ToolType
}

type PackageIndexType struct {
	Url      string
	Notes    string
	Source   SourceType
	Packages []PackageType
}

type PackageProviderType struct {
	Id           string
	PackageIndex PackageIndexType `yaml:"package-index"`
	Packages     []PackageType
}

func main() {
	packageIndexPath := paths.New("e:/git/ino-hardware-package-list/registry.yml")
	rawRegistry, err := packageIndexPath.ReadFile()
	if err != nil {
		panic(err)
	}
	var registryData []PackageProviderType
	err = yaml.Unmarshal(rawRegistry, &registryData)
	if err != nil {
		panic(err)
	}

	for _, packageProvider := range registryData {
		if packageProvider.PackageIndex.Url != "" {
			packageIndexFolderPath, err := paths.TempDir().MkTempDir("inoplatforms-data-generator-package-index-folder")
			defer packageIndexFolderPath.RemoveAll()
			if err != nil {
				panic(err)
			}
			packageIndexPath := packageIndexFolderPath.Join("package_index.json")

			// Download the index data
			httpResponse, err := http.Get(packageProvider.PackageIndex.Url)
			if err != nil {
				slog.Warn("Unable to download package index from %s: %s", err, packageProvider.PackageIndex.Url)
				os.Exit(1)
			}
			defer httpResponse.Body.Close()

			// Write the index data to file
			packageIndexFile, err := packageIndexPath.Create()
			defer packageIndexFile.Close()
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(packageIndexFile, httpResponse.Body); err != nil {
				panic(err)
			}
		}
	}

	// TODO: Unmarshal package index
	// TODO: I need the data type structs, should I use Arduino CLI (either just for the structs, or for loading the index entirely?

	// TODO: Iterate packages and populate the data set

	// TODO: Iterate platforms and populate the data set

	// TODO: Iterate tools and populate the data set

	// TODO: write the site content source files

	// TODO: write the generator data file

	// TODO: marshal the machine readable data file (JSON)

	marshalled, err := yaml.Marshal(registryData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshalled))
}
