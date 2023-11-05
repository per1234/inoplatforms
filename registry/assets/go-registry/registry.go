// Package registry processes the inoplatforms registry.
package registry

import (
	"github.com/arduino/go-paths-helper"
	"gopkg.in/yaml.v3"
)

// Type is the type for data.
type Type struct {
	PackageProviders []PackageProviderType // PackageProviders is the list of providers of packages.
}

// RepositoryType is the type for source repository data.
type RepositoryType struct {
	Url   string // Url is the repository URL.
	Ref   string // Ref is the Git ref (e.g., branch name) in the repository.
	Path  string // Path is the path of the source under the repository.
	Notes string // Notes is supplemental information about the source repository.
}

// SourceType is the type for source data.
type SourceType struct {
	Repository RepositoryType // Repository is data about the source repository.
	Notes      string         // Notes is supplemental information about the source.
}

// PlatformType is the type for platform data.
type PlatformType struct {
	Architecture           string               // Architecture is the machine identifier for the platform.
	Source                 SourceType           // Source contains data about the platform source code.
	InstallationReferences []string             // InstallationReferences contains links to platform installation documentation.
	ToolsDependencies      []ToolDependencyType // ToolsDependencies contains data about the platform's tool dependencies.
	Notes                  string               // Notes is supplemental information about the platform.
}

// ToolDependencyType is the type for tool dependency data.
type ToolDependencyType struct {
	Packager          string // Packager is the machine identifier of the tool dependency's package.
	PackageProviderID string // PackageProviderID is the ID number of the tool dependency's package provider.
}

// ToolType is the type for tool data.
type ToolType struct {
	Name   string     // Name is the machine identifier for the tool.
	Source SourceType // Source contains data about the tool source code.
	Notes  string     // Notes is supplemental information about the tool.
}

type PackageType struct {
	Name      string         // Name is the machine identifier of the package.
	Platforms []PlatformType // Platforms is the platforms provided by the package.
	Tools     []ToolType     // Tools is the tools provided by the package.
	Notes     string         // Notes is supplemental information about the package.
}

// PackageIndexType is the type for package index data.
type PackageIndexType struct {
	Url    string     // Url is the publication URL for the package index.
	Source SourceType // Source is data about the package index source.
	Notes  string     // Notes is supplemental information about the package index.
}

// PackageProviderType is the type for package provider data.
type PackageProviderType struct {
	Id           string           // Id is an arbitrary unique ID number for the package provider.
	PackageIndex PackageIndexType // PackageIndex is data about the package index that provides the packages.
	Packages     []PackageType    // Packages is the list of provided packages.
	Notes        string           // Notes is supplemental information about the package provider.
}

// Load loads data from the registry file.
func Load(registryPath paths.Path) Type {
	// Load
	rawRegistry, err := registryPath.ReadFile()
	if err != nil {
		panic(err)
	}

	// Unmarshal
	var registryData Type
	err = yaml.Unmarshal(rawRegistry, &registryData)
	if err != nil {
		panic(err)
	}

	return registryData
}
