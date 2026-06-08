package cmd

import (
	"fmt"
)

const (
	VERSION = "v0.0.2"
)

// Version prints the version of the program, wow.
func Version() {
	fmt.Printf("agopass %s :: Author © 2026 Alvesafk\n", VERSION)
}
