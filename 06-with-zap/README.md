# 06-with-zap

Emit Zap log lines tagged with build metadata, in both flat (`Fields()`) and grouped (`Namespace()`) shapes.

## Run

```sh
go run .
```

## Expected output

```json
{"level":"info","msg":"flat fields — build_* keys appear at top level","build_version":"dev","build_commit":"...","build_branch":"unknown","build_time":"...","build_goversion":"go1.24.x"}
{"level":"info","msg":"grouped namespace — build:{...} sub-object","build":{"version":"dev","commit":"...","branch":"unknown","time":"...","goversion":"go1.24.x"}}
```

(Timestamps are suppressed in the example for clean diffing.)

## When to pick which

| Use | When |
|-----|------|
| `Fields()` | Log-aggregation pipelines that expect flat `build_*` keys (Datadog, Loki, ELK with default mappings). |
| `Namespace()` | When you want a single `build` sub-object — keeps the rest of the log line uncluttered. |

## Real-world: production zap config

In a production service you typically tag the *root* logger so every
log line carries the build metadata, then derive child loggers from it:

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    bizap "github.com/ubgo/buildinfo/contrib/buildinfo-zap"
)

prodCfg := zap.NewProductionConfig()
prodCfg.EncoderConfig.TimeKey = "ts"
prodCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

base, err := prodCfg.Build()
if err != nil { panic(err) }

// Tag every log line with the namespaced build sub-object.
logger := base.With(bizap.Namespace())
defer logger.Sync()

logger.Info("startup")
// → {"level":"info","ts":"2026-…","msg":"startup","build":{"version":"1.2.3",...}}
```

`With` is called once at boot and the resulting logger is passed
through your service. Every subsequent `.Info`/`.Error`/etc. carries
the `build` group automatically — no per-call boilerplate.

## Combining with `-ldflags` stamping

Build with `-ldflags` populated values (see [`02-ldflags-stamping`](../02-ldflags-stamping))
and the build group in your logs becomes meaningful:

```json
{"level":"info","msg":"startup","build":{"version":"1.2.3","commit":"abc1234","branch":"main",...}}
```

Now every log line is correlatable to a specific deploy — useful when
filtering Sentry / Loki / ELK for "errors from version 1.2.3".
