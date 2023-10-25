// Package search is for the site search.
package search

import "github.com/arduino/go-paths-helper"

// Type is the type for search data.
type Type struct {
	PlatformDocuments PlatformDocuments
	ToolDocuments     ToolDocuments
}

// PlatformDocuments is the type for platforms search data.
type PlatformDocuments []PlatformDocument

// PlatformDocument is the type for individual platform search data.
type PlatformDocument struct {
	Ref          string
	Content      string
	Vendor       string
	Architecture string
}

// ToolDocuments is the type for tool search data.
type ToolDocuments []ToolDocument

// ToolDocument is the type for individual tool search data.
type ToolDocument struct {
}

// Write creates the search data file.
func (data Type) Write(dataFolder paths.Path) {}
