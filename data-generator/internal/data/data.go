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
	InstallationReferences []string
	ToolsDependencies      []ToolDependencyType
	DiscoveryDependencies  []ToolType
	MonitorDependencies    []ToolType
	Releases               []PlatformReleaseType
	// TODO: SearchData SearchDataType
}

// TODO:
// type SearchDataType struct{
// Ref string
// SearchContent string
// }

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
	ToolsDependencies     []ToolDependencyType
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
	// TODO: SearchData SearchDataType
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
	PackageIndex PackageIndexType
	Packages     []PackageType
	Status       StatusType
}

// StatusType is the type for metadata about the status of the data.
type StatusType struct {
	Result  ResultType
	Message string
}

type ResultType int

const (
	Success ResultType = iota
	Failure
)

type Type struct {
	PackageProviders []PackageProviderType
	Stats            StatsType
}

type StatsType struct {
	PlatformCount     int
	ArchitectureCount int
	Timestamp         string
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
