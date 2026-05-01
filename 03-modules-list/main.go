// Example 03-modules-list — dump the dependency module list (SBOM-friendly).
//
// runtime/debug.ReadBuildInfo populates Info.Modules with every direct and
// indirect dependency module the build linked, including replacements applied
// via go.mod replace directives.
package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/ubgo/buildinfo"
)

func main() {
	mods := buildinfo.Get().Modules
	fmt.Printf("--- %d dependency modules ---\n\n", len(mods))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "PATH\tVERSION\tSUM")
	for _, m := range mods {
		fmt.Fprintf(w, "%s\t%s\t%s\n", m.Path, m.Version, m.Sum)
	}
	_ = w.Flush()
}
