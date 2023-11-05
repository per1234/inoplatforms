// Package data provides code for working with the inoplatforms catalog data.
package data

// Type is the type for data.
type Type struct {
	PackageProviders []PackageProviderType `json:"package_providers"` // PackageProviders is the list of providers of packages.
	Stats            StatsType             `json:"stats"`             // Stats is statistics about the data overall.
}

// RepositoryType is the type for source repository data.
type RepositoryType struct {
	Url   string `json:"url"`   // Url is the repository URL.
	Ref   string `json:"ref"`   // Ref is the Git ref (e.g., branch name) in the repository.
	Path  string `json:"path"`  // Path is the path of the source under the repository.
	Notes string `json:"notes"` // Notes is supplemental information about the source repository.
}

// SourceType is the type for source data.
type SourceType struct {
	Repository RepositoryType `json:"repository"` // Repository is data about the source repository.
	Notes      string         `json:"notes"`      // Notes is supplemental information about the source.
}

// ToolDependencyType is the type for tool dependency data.
type ToolDependencyType struct {
	Packager          string `json:"packager"`            // Packager is the machine identifier of the tool dependency's package.
	Name              string `json:"name"`                // Name is the machine identifier of the tool dependency.
	Version           string `json:"version"`             // Version is the version number of the tool dependency.
	PackageProviderID string `json:"package_provider_id"` // PackageProviderID is the ID number of the tool dependency's package provider.
}

// PlatformType is the type for platform data.
type PlatformType struct {
	Name                   string                `json:"name"`                    // Name is the human identifier for the latest release of the platform.
	Architecture           string                `json:"architecture"`            // Architecture is the machine identifier for the platform.
	Deprecated             bool                  `json:"deprecated"`              // Deprecated indicates whether the platform is deprecated.
	Help                   string                `json:"help"`                    // Help contains data about user support for the platform.
	Boards                 []BoardType           `json:"boards"`                  // Boards contains data about the boards supported by the latest release of the platform.
	Source                 SourceType            `json:"source"`                  // Source contains data about the platform source code.
	InstallationReferences []string              `json:"installation_references"` // InstallationReferences contains links to platform installation documentation.
	ToolsDependencies      []ToolDependencyType  `json:"tools_dependencies"`      // ToolsDependencies is the standard tool dependencies of the latest release of the platform.
	DiscoveryDependencies  []ToolType            `json:"discovery_dependencies"`  // DiscoveryDependencies is the pluggable discovery tool dependencies of latest release of the platform.
	MonitorDependencies    []ToolType            `json:"monitor_dependencies"`    // MonitorDependencies is the pluggable monitor tool dependencies of the latest release of the platform.
	Releases               []PlatformReleaseType `json:"releases"`                // Releases contains data about each of the releases of the platform.
	Notes                  string                `json:"notes"`                   // Notes is supplemental information about the platform.
	// TODO: SearchData SearchDataType `json:"search_data"`	// SearchData contains site search index data for the platform.
}

// TODO:
// // SearchDataType is the type for site search index data.
// type SearchDataType struct{
// Ref string `json:"ref"`	// Ref is the search result entry displayed to the visitor.
// SearchContent string `json:"search_content"`	// SearchContent is the text content for keyword search.
// }

// BoardType is the type for data about a board supported by the platform.
type BoardType struct {
	Name string `json:"name"` // Name is the human identifier for the board.
}

// PlatformReleaseType is the type for platform release data.
type PlatformReleaseType struct {
	Name                  string               `json:"name"`                   // Name is the human identifier for the platform.
	Version               string               `json:"version"`                // Version is the platform release version number.
	Help                  string               `json:"help"`                   // Help contains data about user support for the platform release.
	ArchiveUrl            string               `json:"archive_url"`            // ArchiveUrl is the platform release archive download URL.
	Checksum              string               `json:"checksum"`               // Checksum is the platform release archive checksum.
	Boards                []BoardType          `json:"boards"`                 // Boards contains data about the boards supported by the platform release.
	ToolsDependencies     []ToolDependencyType `json:"tools_dependencies"`     // ToolsDependencies is the standard tool dependencies of the platform release.
	DiscoveryDependencies []ToolType           `json:"discovery_dependencies"` // DiscoveryDependencies is the pluggable discovery tool dependencies of the platform release.
	MonitorDependencies   []ToolType           `json:"monitor_dependencies"`   // MonitorDependencies is the pluggable monitor tool dependencies of the platform release.
}

// SystemType is the type for host-specific tool release data.
type SystemType struct {
	Checksum   string `json:"checksum"`    // Checksum is the tool release archive checksum.
	Host       string `json:"host"`        // Host is the host architecture machine identifier.
	ArchiveUrl string `json:"archive_url"` // ArchiveUrl is the tool release archive download URL.
}

// ToolType is the type for tool data.
type ToolType struct {
	Name     string            `json:"name"`     // Name is the machine identifier for the tool.
	Source   SourceType        `json:"source"`   // Source contains data about the tool source code.
	Releases []ToolReleaseType `json:"releases"` // Releases contains data about each of the releases of the tool.
	Notes    string            `json:"notes"`    // Notes is supplemental information about the tool.
	// TODO: SearchData SearchDataType	 `json:"search_data"` // SearchData contains site search index data for the tool.
}

type ToolReleaseType struct {
	Version string       `json:"version"` // Version is the tool release version.
	Systems []SystemType `json:"systems"` // Systems contains target host-specific data.
}

type PackageType struct {
	Name       string         `json:"name"`           // Name is the machine identifier of the package.
	Maintainer string         `json:"maintainer_url"` // Maintainer is the human identifier for the maintainer of the package.
	WebsiteUrl string         `json:"website_url"`    // WebsiteURL is the URL for information about the package.
	Platforms  []PlatformType `json:"platforms"`      // Platforms is the platforms provided by the package.
	Tools      []ToolType     `json:"tools"`          // Tools is the tools provided by the package.
	Notes      string         `json:"notes"`          // Notes is supplemental information about the package.
}

// PackageIndexType is the type for package index data.
type PackageIndexType struct {
	Url    string     `json:"url"`    // Url is the publication URL for the package index.
	Source SourceType `json:"source"` // Source is data about the package index source.
	Notes  string     `json:"notes"`  // Notes is supplemental information about the package index.
}

// PackageProviderType is the type for package provider data.
type PackageProviderType struct {
	Id           string           `json:"id"`            // Id is an arbitrary unique ID number for the package provider.
	PackageIndex PackageIndexType `json:"package_index"` // PackageIndex is data about the package index that provides the packages.
	Packages     []PackageType    `json:"packages"`      // Packages is the list of provided packages.
	Notes        string           `json:"notes"`         // Notes is supplemental information about the package provider.
	Status       StatusType       `json:"status"`        // Status contains data about the status of the collected data for the package provider.
}

// StatusType is the type for metadata about the status of the collected data.
type StatusType struct {
	Result  ResultType `json:"result"`  // Result is the data collection result.
	Message string     `json:"message"` // Message is a message for humans about the status of the collected data.
}

// ResultType is the type for the data collection result.
type ResultType int

const (
	Success ResultType = iota // Success indicates data collection was successful.
	Failure                   // Failure indicates data collection failed.
)

// StatsType is the type for data statistics.
type StatsType struct {
	PlatformCount     int    `json:"platform_count"`     // PlatformCount is the number of platforms.
	ArchitectureCount int    `json:"architecture_count"` // ArchitectureCount is the number of unique architectures.
	Timestamp         string `json:"timestamp"`          // Timestamp is the timestamp of when the data was generated.
}
