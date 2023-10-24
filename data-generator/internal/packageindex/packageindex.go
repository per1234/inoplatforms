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

// Type is the type for package index data.
type Type struct {
	Packages []*PackageType // Packages is the packages of the package index.
}

// PackageType is the type for package data.
type PackageType struct {
	Name       string                 // Name is the machine identifier of the package.
	Maintainer string                 // Maintainer is the human identifier for the maintainer of the package.
	WebsiteURL string                 // WebsiteURL is the URL for information about the package.
	URL        string                 // URL is the URL for information about the package.
	Email      string                 // Email is the email address to contact the package maintainer.
	Platforms  []*PlatformReleaseType // Platforms is the platform releases provided by the package.
	Tools      []*ToolReleaseType     // Tools is the tool releases provided by the package.
	Help       HelpType               // Help contains data about user support for the package.
}

// PlatformReleaseType is the type for platform release data.
type PlatformReleaseType struct {
	Name                  string                           // Name is the human identifier for the platform.
	Architecture          string                           // Architecture is the machine identifier for the platform.
	Version               *semver.Version                  // Version is the platform release version number.
	Deprecated            bool                             // Deprecated indicates whether the platform is deprecated.
	Category              string                           // Category is the platform's category.
	URL                   string                           // URL is the platform release archive download URL.
	ArchiveFileName       string                           // ArchiveFileName is the platform release archive filename.
	Checksum              string                           // Checksum is the platform release archive checksum.
	Size                  json.Number                      // Size is the platform release archive file size.
	Boards                []BoardType                      // Boards contains data about the boards supported by the platform release.
	Help                  HelpType                         // Help contains data about user support for the platform release.
	ToolsDependencies     []VersionedToolDependencyType    // ToolsDependencies is the standard tool dependencies of the platform release.
	DiscoveryDependencies []NonVersionedToolDependencyType // DiscoveryDependencies is the pluggable discovery tool dependencies of the platform release.
	MonitorDependencies   []NonVersionedToolDependencyType // MonitorDependencies is the pluggable monitor tool dependencies of the platform release.
}

// NonVersionedToolDependencyType is the type for versioned tool dependency data.
type VersionedToolDependencyType struct {
	Packager string
	Name     string
	Version  *semver.RelaxedVersion
}

// NonVersionedToolDependencyType is the type for non-versioned tool dependency data.
type NonVersionedToolDependencyType struct {
	Packager string // Packager is the vendor name of the tool.
	Name     string // Name is the name of the tool.
}

// ToolReleaseType is the type for tool release data.
type ToolReleaseType struct {
	Name    string                   // Name is the machine identifier for the tool.
	Version *semver.RelaxedVersion   // Version is the tool release version.
	Systems []ToolReleaseFlavourType // Systems contains target host-specific data.
}

// ToolReleaseFlavourType is the type for host-specific tool release data.
type ToolReleaseFlavourType struct {
	OS              string      // OS is the host architecture machine identifier.
	URL             string      // URL is the tool release archive download URL.
	ArchiveFileName string      // ArchiveFileName is the tool release archive filename.
	Size            json.Number // Size is the tool release archive file size.
	Checksum        string      // Checksum is the tool release archive checksum.
}

// BoardType is the type for data about a board supported by the platform.
type BoardType struct {
	Name string        // Name is the human identifier for the board.
	ID   []BoardIDType // ID contains identification data for the board.
}

// BoardIDType is the type for board identification data.
type BoardIDType struct {
	USB string // USB is the USB identification for the board.
}

// HelpType is the type for user support request data.
type HelpType struct {
	Online string // Online is the online support request URL.
}

// Get downloads and parses a package index.
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
	rawIndex, err := downloadPath.ReadFile()
	if err != nil {
		return Type{}, err
	}
	var indexData Type
	err = json.Unmarshal(rawIndex, &indexData)
	if err != nil {
		return Type{}, err
	}

	return indexData, nil
}
