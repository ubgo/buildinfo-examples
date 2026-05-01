// Example 06-with-zap — emit Zap log lines tagged with build metadata,
// shown in both flat (Fields) and grouped (Namespace) shapes.
package main

import (
	"go.uber.org/zap"

	bzap "github.com/ubgo/buildinfo/contrib/buildinfo-zap"
)

func main() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "" // hide timestamps so the output diffs cleanly in tests
	logger, _ := cfg.Build()
	defer func() { _ = logger.Sync() }()

	flat := logger.With(bzap.Fields()...)
	flat.Info("flat fields — build_* keys appear at top level")

	grouped := logger.With(bzap.Namespace())
	grouped.Info("grouped namespace — build:{...} sub-object")
}
