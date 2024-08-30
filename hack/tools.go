// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package main

import (
	// Register code-generator.
	_ "k8s.io/code-generator"
)
