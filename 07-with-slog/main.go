// Example 07-with-slog — emit stdlib slog log lines tagged with build metadata,
// shown in both flat (Attrs) and grouped (Group) shapes.
package main

import (
	"log/slog"
	"os"

	bslog "github.com/ubgo/buildinfo/contrib/buildinfo-slog"
)

func main() {
	// Hide the auto-generated time key so output diffs cleanly in tests.
	opts := &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)

	flat := slog.New(handler)
	for _, a := range bslog.Attrs() {
		flat = flat.With(a)
	}
	flat.Info("flat attrs — build_* keys appear at top level")

	grouped := slog.New(handler).With(bslog.Group())
	grouped.Info("grouped — build:{...} sub-object")
}
