// +build tools

// Package tools adds references to tools like linters as dependencies.
// It will not actually be included in the regular build.
package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
