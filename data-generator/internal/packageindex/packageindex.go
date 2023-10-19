// Package packageindex processes Arduino package indexes.
package packageindex

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/arduino/go-paths-helper"
	semver "go.bug.st/relaxed-semver"
	"golang.org/x/exp/slog"
)

type Type struct {
	Packages []*Package `json:"packages"`
}

// indexPackage represents a single entry from package_index.json file.
//
//easyjson:json
type Package struct {
	Name       string             `json:"name"`
	Maintainer string             `json:"maintainer"`
	WebsiteURL string             `json:"websiteUrl"`
	URL        string             `json:"Url"`
	Email      string             `json:"email"`
	Platforms  []*PlatformRelease `json:"platforms"`
	Tools      []*ToolRelease     `json:"tools"`
	Help       Help               `json:"help,omitempty"`
}

// PlatformRelease represents a single Core Platform from package_index.json file.
//
//easyjson:json
type PlatformRelease struct {
	Name                  string                `json:"name"`
	Architecture          string                `json:"architecture"`
	Version               *semver.Version       `json:"version"`
	Deprecated            bool                  `json:"deprecated"`
	Category              string                `json:"category"`
	URL                   string                `json:"url"`
	ArchiveFileName       string                `json:"archiveFileName"`
	Checksum              string                `json:"checksum"`
	Size                  json.Number           `json:"size"`
	Boards                []Board               `json:"boards"`
	Help                  Help                  `json:"help,omitempty"`
	ToolDependencies      []ToolDependency      `json:"toolsDependencies"`
	DiscoveryDependencies []DiscoveryDependency `json:"discoveryDependencies"`
	MonitorDependencies   []MonitorDependency   `json:"monitorDependencies"`
}

// indexToolDependency represents a single dependency of a core from a tool.
//
//easyjson:json
type ToolDependency struct {
	Packager string                 `json:"packager"`
	Name     string                 `json:"name"`
	Version  *semver.RelaxedVersion `json:"version"`
}

// indexDiscoveryDependency represents a single dependency of a core from a pluggable discovery tool.
//
//easyjson:json
type DiscoveryDependency struct {
	Packager string `json:"packager"`
	Name     string `json:"name"`
}

// indexMonitorDependency represents a single dependency of a core from a pluggable monitor tool.
//
//easyjson:json
type MonitorDependency struct {
	Packager string `json:"packager"`
	Name     string `json:"name"`
}

// indexToolRelease represents a single Tool from package_index.json file.
//
//easyjson:json
type ToolRelease struct {
	Name    string                 `json:"name"`
	Version *semver.RelaxedVersion `json:"version"`
	Systems []ToolReleaseFlavour   `json:"systems"`
}

// indexToolReleaseFlavour represents a single tool flavor in the package_index.json file.
//
//easyjson:json
type ToolReleaseFlavour struct {
	OS              string      `json:"host"`
	URL             string      `json:"url"`
	ArchiveFileName string      `json:"archiveFileName"`
	Size            json.Number `json:"size"`
	Checksum        string      `json:"checksum"`
}

// indexBoard represents a single Board as written in package_index.json file.
//
//easyjson:json
type Board struct {
	Name string    `json:"name"`
	ID   []BoardID `json:"id,omitempty"`
}

// indexBoardID represents the ID of a single board. i.e. uno, yun, diecimila, micro and the likes
//
//easyjson:json
type BoardID struct {
	USB string `json:"usb"`
}

// indexHelp represents the help URL
//
//easyjson:json
type Help struct {
	Online string `json:"online,omitempty"`
}

func Get(url string) (Type, error) {
	// Download the index.
	httpResponse, err := http.Get(url)
	if err != nil {
		slog.Warn("Unable to download package index from %s: %s", err, url)
		os.Exit(1)
	}
	defer httpResponse.Body.Close()

	// Write the index data to file
	downloadFolderPath, err := paths.TempDir().MkTempDir("inoplatforms-data-generator-package-index-folder")
	defer downloadFolderPath.RemoveAll()
	if err != nil {
		panic(err)
	}
	downloadPath := downloadFolderPath.Join("package_index.json")

	indexFile, err := downloadPath.Create()
	defer indexFile.Close()
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(indexFile, httpResponse.Body); err != nil {
		panic(err)
	}

	// Unmarshal
	rawIndex, err := indexFile.ReadFile()
	if err != nil {
		return nil, err
	}
	var indexData Type
	err = json.Unmarshal(rawIndex, &indexData)
	if err != nil {
		return nil, err
	}

	return indexData, nil
}
