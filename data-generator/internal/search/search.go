// Package search is for the site search.
package search

import "github.com/arduino/go-paths-helper"

type Type struct {
	PlatformDocuments PlatformDocuments
	ToolDocuments     ToolDocuments
}

type PlatformDocuments []PlatformDocument

type PlatformDocument struct {
	Ref          string
	Content      string
	Vendor       string
	Architecture string
}

type ToolDocuments []ToolDocument

type ToolDocument struct {
}

func (data Type) Write(dataFolder paths.Path) {}
