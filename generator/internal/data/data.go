// Package data works with the generated data.
package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/arduino/go-paths-helper"
	"github.com/per1234/inoplatforms/generator/data"
	"github.com/per1234/inoplatforms/generator/internal/packageindex"
	"github.com/per1234/inoplatforms/registry/assets/go-registry/registry"
)

// Get collects data.
func Get(registry registry.Type) data.Type {
	var data data.Type
	for _, packageProvider := range registry.PackageProviders {
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

	return data
}

// marshal returns the data marshaled into JSON format in byte encoding.
func marshal(data data.Type) []byte {
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
func Write(data data.Type, dataFilePath paths.Path) error {
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

	err = dataFilePath.WriteFile(marshal(data))
	if err != nil {
		return fmt.Errorf("While writing data file: %v", err)
	}

	return nil
}
