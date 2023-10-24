// Package data works with the generated data.
package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/arduino/go-paths-helper"
	"github.com/per1234/inoplatforms/data-generator/internal/packageindex"
)

// RepositoryType is the type for source repository data.
type RepositoryType struct {
	Url  string // Url is the repository URL.
	Ref  string // Ref is the Git ref (e.g., branch name) in the repository.
	Path string // Path is the path of the source under the repository.
}

// SourceType is the type for source data.
type SourceType struct {
	Repository RepositoryType // Repository is data about the source repository.
}

// ToolDependencyType is the type for tool dependency data.
type ToolDependencyType struct {
	Packager          string // Packager is the machine identifier of the tool dependency's package.
	Name              string // Name is the machine identifier of the tool dependency.
	Version           string // Version is the version number of the tool dependency.
	PackageProviderID string // PackageProviderID is the ID number of the tool dependency's package provider.
}

// PlatformType is the type for platform data.
type PlatformType struct {
	Name                   string                // Name is the human identifier for the latest release of the platform.
	Architecture           string                // Architecture is the machine identifier for the platform.
	Deprecated             bool                  // Deprecated indicates whether the platform is deprecated.
	Help                   string                // Help contains data about user support for the platform.
	Boards                 []BoardType           // Boards contains data about the boards supported by the latest release of the platform.
	Source                 SourceType            // Source contains data about the platform source code.
	InstallationReferences []string              // InstallationReferences contains links to platform installation documentation.
	ToolsDependencies      []ToolDependencyType  // ToolsDependencies is the standard tool dependencies of the latest release of the platform.
	DiscoveryDependencies  []ToolType            // DiscoveryDependencies is the pluggable discovery tool dependencies of latest release of the platform.
	MonitorDependencies    []ToolType            // MonitorDependencies is the pluggable monitor tool dependencies of the latest release of the platform.
	Releases               []PlatformReleaseType // Releases contains data about each of the releases of the platform.
	// TODO: SearchData SearchDataType
}

// TODO:
// // SearchDataType is the type for site search index data.
// type SearchDataType struct{
// Ref string
// SearchContent string
// }

// BoardType is the type for data about a board supported by the platform.
type BoardType struct {
	Name string // Name is the human identifier for the board.
}

// PlatformReleaseType is the type for platform release data.
type PlatformReleaseType struct {
	Name                  string               // Name is the human identifier for the platform.
	Version               string               // Version is the platform release version number.
	Help                  string               // Help contains data about user support for the platform release.
	ArchiveUrl            string               // ArchiveUrl is the platform release archive download URL.
	Checksum              string               // Checksum is the platform release archive checksum.
	Boards                []BoardType          // Boards contains data about the boards supported by the platform release.
	ToolsDependencies     []ToolDependencyType // ToolsDependencies is the standard tool dependencies of the platform release.
	DiscoveryDependencies []ToolType           // DiscoveryDependencies is the pluggable discovery tool dependencies of the platform release.
	MonitorDependencies   []ToolType           // MonitorDependencies is the pluggable monitor tool dependencies of the platform release.
}

// SystemType is the type for host-specific tool release data.
type SystemType struct {
	Checksum   string // Checksum is the tool release archive checksum.
	Host       string // Host is the host architecture machine identifier.
	ArchiveUrl string // ArchiveUrl is the tool release archive download URL.
}

// ToolType is the type for tool data.
type ToolType struct {
	Name     string            // Name is the machine identifier for the tool.
	Source   SourceType        // Source contains data about the tool source code.
	Releases []ToolReleaseType // Releases contains data about each of the releases of the tool.
	// TODO: SearchData SearchDataType	// SearchData contains site search index data for the tool.
}

type ToolReleaseType struct {
	Version string       // Version is the tool release version.
	Systems []SystemType // Systems contains target host-specific data.
}

type PackageType struct {
	Name       string         // Name is the machine identifier of the package.
	Maintainer string         // Maintainer is the human identifier for the maintainer of the package.
	WebsiteUrl string         // WebsiteURL is the URL for information about the package.
	Platforms  []PlatformType // Platforms is the platforms provided by the package.
	Tools      []ToolType     // Tools is the tools provided by the package.
}

// PackageIndexType is the type for package index data.
type PackageIndexType struct {
	Url    string     // Url is the publication URL for the package index.
	Notes  string     // Notes is notes about the package index.
	Source SourceType // Source is data about the package index source.
}

// PackageProviderType is the type for package provider data.
type PackageProviderType struct {
	Id           string           // Id is an arbitrary unique ID number for the package provider.
	PackageIndex PackageIndexType // PackageIndex is data about the package index that provides the packages.
	Packages     []PackageType    // Packages is the list of provided packages.
	Status       StatusType       // Status contains data about the status of the collected data for the package provider.
}

// StatusType is the type for metadata about the status of the collected data.
type StatusType struct {
	Result  ResultType // Result is the data collection result.
	Message string     // Message is a message for humans about the status of the collected data.
}

// ResultType is the type for the data collection result.
type ResultType int

const (
	Success ResultType = iota // Success indicates data collection was successful.
	Failure                   // Failure indicates data collection failed.
)

// Type is the type for data.
type Type struct {
	PackageProviders []PackageProviderType // PackageProviders is the list of providers of packages.
	Stats            StatsType             // Stats is statistics about the data overall.
}

// StatsType is the type for data statistics.
type StatsType struct {
	PlatformCount     int    // PlatformCount is the number of platforms.
	ArchitectureCount int    // ArchitectureCount is the number of unique architectures.
	Timestamp         string // Timestamp is the timestamp of when the data was generated.
}

// Populate gathers all additional data.
func (data *Type) Populate() {
	for _, packageProvider := range data.PackageProviders {
		if packageProvider.PackageIndex.Url != "" {
			packageindex.Get(packageProvider.PackageIndex.Url)
			// TODO: Unmarshal package index
			// TODO: I need the data type structs, should I use Arduino CLI (either just for the structs, or for loading the index entirely?

			// TODO: Iterate packages and populate the data set

			// TODO: Iterate platforms and:
			// - populate the data set
			// - populate platform search index data

			// TODO: update platform count and architecture count in data.Stats

			// TODO: Iterate tools and:
			// - populate the data set
			// - populate tool search index data
		}
	}

	timestamp := time.Now().UTC()
	data.Stats.Timestamp = timestamp.Format("2006-01-02 15:04:05 UTC")
}

// marshal returns the data marshaled into JSON format in byte encoding.
func (data Type) marshal() []byte {
	var marshaledDataBuffer bytes.Buffer

	jsonEncoder := json.NewEncoder(io.Writer(&marshaledDataBuffer))

	// By default, the json package HTML-sanitizes strings during marshaling (https://golang.org/pkg/encoding/json/#Marshal)
	// This means that the simple json.MarshalIndent() approach would result in the report containing gibberish.
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "  ")
	err := jsonEncoder.Encode(data)
	if err != nil {
		panic(fmt.Errorf("Error while marshalling generated data: %v", err))
	}

	return marshaledDataBuffer.Bytes()
}

// Write writes the data to a JSON file created at the specified path.
func (data Type) Write(dataFilePath paths.Path) error {
	dataFilePathParentExists, err := dataFilePath.Parent().ExistCheck()
	if err != nil {
		return fmt.Errorf("Problem processing data file output path %v: %v", dataFilePath, err)
	}
	if !dataFilePathParentExists {
		err = dataFilePath.Parent().MkdirAll()
		if err != nil {
			return fmt.Errorf("Unable to create data file path (%v): %v", dataFilePath.Parent(), err)
		}
	}

	err = dataFilePath.WriteFile(data.marshal())
	if err != nil {
		return fmt.Errorf("While writing data file: %v", err)
	}

	return nil
}
