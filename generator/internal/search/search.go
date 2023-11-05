// Package search is for the site search.
package search

import "github.com/arduino/go-paths-helper"

// Type is the type for search data.
type Type struct {
	PlatformDocuments PlatformDocuments // PlatformDocuments is the platforms search data.
	ToolDocuments     ToolDocuments     // ToolDocuments is the tools search data.
}

// PlatformDocuments is the type for platforms search data.
type PlatformDocuments []PlatformDocument

// PlatformDocument is the type for individual platform search data.
type PlatformDocument struct {
	Ref          string // Ref is the presentation of the search result to the user.
	Content      string // Content is the free text for keyword search.
	Packager     string // Packager is the machine identifier of the platform's package.
	Architecture string // Architecture is the machine identifier of the platform within its package.
}

// ToolDocuments is the type for tool search data.
type ToolDocuments []ToolDocument

// ToolDocument is the type for individual tool search data.
type ToolDocument struct {
	Ref      string // Ref is the presentation of the search result to the user.
	Content  string // Content is the free text for keyword search.
	Packager string // Packager is the machine identifier of the tool's package.
	Name     string // Name is the machine identifier of the tool within its package.
}

// Write creates the search data file.
func (data Type) Write(dataFolder paths.Path) {}
