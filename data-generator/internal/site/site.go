// Package site generates the website source content.
package site

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/arduino/go-paths-helper"
	"github.com/per1234/inoplatforms/data-generator/internal/data"
)

// TODO: Add frontmatter fields to page types:
// PackagesType is the type for packages page data.
type PackagesType struct{}

// PlatformType is the type for platform page data.
type PlatformType struct{}

// PlatformReleaseType is the type for platform release page data.
type PlatformReleaseType struct{}

// ToolType is the type for tool page data.
type ToolType struct{}

// ToolReleaseType is the type for tool release page data.
type ToolReleaseType struct{}

// WritePages writes the page content source files.
func WritePages(data data.Type, siteContentPath paths.Path) {
	for _, packageProvider := range data.PackageProviders {
		// Extract page front matter data into page structs.
		// TODO

		// Marshal front matter.
		// TODO

		// Write page source file.
		// TODO
	}
}

// marshal returns the data marshaled into JSON format in byte encoding.
func marshalGeneratorData(data data.StatsType) []byte {
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

// WriteGeneratorData writes the data to a JSON file created at the specified path.
func WriteGeneratorData(data data.Type, generatorDataPath paths.Path) err {
	generatorDataPathParentExists, err := generatorDataPath.Parent().ExistCheck()
	if err != nil {
		return fmt.Errorf("Problem processing generator data file output path %v: %v", generatorDataPath, err)
	}
	if !generatorDataPathParentExists {
		err = generatorDataPath.Parent().MkdirAll()
		if err != nil {
			return fmt.Errorf("Unable to create data file path (%v): %v", generatorDataPath.Parent(), err)
		}
	}

	err = generatorDataPath.WriteFile(data.Stats.marshal())
	if err != nil {
		return fmt.Errorf("While writing data file: %v", err)
	}

	return nil
}
