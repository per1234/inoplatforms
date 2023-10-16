//go:build tools

// Package tools defines the project's Go module tool dependencies.
// See: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/editorconfig-checker/editorconfig-checker/cmd/editorconfig-checker"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)