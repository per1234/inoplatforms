package main

import (
	"fmt"

	"github.com/arduino/go-paths-helper"
	"gopkg.in/yaml.v3"
)

type RepositoryType struct {
	Url    string
	Branch string
	Path   string
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

	marshalled, err := yaml.Marshal(registryData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshalled))
}
