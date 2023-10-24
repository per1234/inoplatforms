// Package registry processes the inoplatforms registry.
package registry

import (
	"github.com/arduino/go-paths-helper"
	"github.com/per1234/inoplatforms/data-generator/internal/data"
	"gopkg.in/yaml.v2"
)

// Load loads data from the registry file.
func Load(registryPath paths.Path) data.Type {
	// Load
	rawRegistry, err := registryPath.ReadFile()
	if err != nil {
		panic(err)
	}

	// Unmarshal
	var registryData data.Type
	err = yaml.Unmarshal(rawRegistry, &registryData)
	if err != nil {
		panic(err)
	}

	return registryData
}
