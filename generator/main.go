package main

import (
	"fmt"
	"log/slog"

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
	Packager string
	Id       string
}

type PlatformType struct {
	Architecture           string
	Source                 SourceType
	InstallationReferences []string             `yaml:"installation-references"`
	ToolsDependencies      []ToolDependencyType `yaml:"tools-dependencies"`
}

type PackageType struct {
	Name      string
	Platforms []PlatformType
}

type PackageIndexType struct {
	Url      string
	Notes    string
	Source   SourceType
	Packages []PackageType
}

type RegistrationType struct {
	Id           string
	PackageIndex PackageIndexType `yaml:"package-index"`
	Packages     []PackageType
}

func main() {
	slog.Info("hello, world")
	packageIndexPath := paths.New("e:/git/ino-hardware-package-list/registry.yml")
	rawRegistry, err := packageIndexPath.ReadFile()
	if err != nil {
		panic(err)
	}
	var registryData []RegistrationType
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
