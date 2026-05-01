// Example 02-ldflags-stamping — CI-style version stamping.
//
// Build with the included Makefile to inject Version / Commit / BuildTime /
// Branch via -ldflags. ldflags overrides win over runtime/debug data.
package main

import (
	"fmt"

	"github.com/ubgo/buildinfo"
)

func main() {
	info := buildinfo.Get()
	fmt.Printf("version=%s commit=%s build_time=%s branch=%s\n",
		info.Version, info.Commit, info.BuildTime, info.Branch)
}
