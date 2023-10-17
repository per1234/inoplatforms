//go:build tools

// Package tools defines the project's Go module tool dependencies.
// See: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/editorconfig-checker/editorconfig-checker/cmd/editorconfig-checker"
	_ "github.com/go-task/task/v3/cmd/task"
	_ "github.com/gohugoio/hugo"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
