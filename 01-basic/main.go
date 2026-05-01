// Example 01-basic — print buildinfo.Get() to stdout.
//
// Demonstrates the zero-config path: buildinfo auto-populates VCS metadata
// from runtime/debug.ReadBuildInfo (Go 1.18+ embeds vcs.revision, vcs.time,
// and vcs.modified) without any -ldflags arguments at build time.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ubgo/buildinfo"
)

func main() {
	info := buildinfo.Get()

	fmt.Println("--- flat fields ---")
	fmt.Printf("version:     %s\n", info.Version)
	fmt.Printf("commit:      %s\n", info.Commit)
	fmt.Printf("build_time:  %s\n", info.BuildTime)
	fmt.Printf("branch:      %s\n", info.Branch)
	fmt.Printf("go_version:  %s\n", info.GoVersion)
	fmt.Printf("goos/goarch: %s/%s\n", info.GOOS, info.GOARCH)
	fmt.Printf("modified:    %t\n", info.Modified)
	fmt.Printf("modules:     %d entries\n", len(info.Modules))

	fmt.Println("\n--- JSON ---")
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(info); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
