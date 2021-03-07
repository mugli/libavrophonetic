// +build buildtools

// Package buildtools adds references to buildtools like linters as dependencies.
// It will not actually be included in the regular build.
package buildtools

import (
	_ "github.com/daixiang0/gci"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
