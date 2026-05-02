# 07-with-slog

Emit stdlib `log/slog` log lines tagged with build metadata, in both flat (`Attrs()`) and grouped (`Group()`) shapes. Zero third-party dependencies.

## Run

```sh
go run .
```

## Expected output

```json
{"level":"INFO","msg":"flat attrs — build_* keys appear at top level","build_version":"dev","build_commit":"...","build_branch":"unknown","build_time":"...","build_goversion":"go1.24.x"}
{"level":"INFO","msg":"grouped — build:{...} sub-object","build":{"version":"dev","commit":"...","branch":"unknown","time":"...","goversion":"go1.24.x"}}
```

(Timestamps are suppressed in the example for clean diffing.)

## When to pick which

| Use | When |
|-----|------|
| `Attrs()` | Log-aggregation pipelines that expect flat `build_*` keys (Datadog, Loki, ELK with default mappings). |
| `Group()` | When you want a single `build` sub-object — keeps the rest of the log line uncluttered. |

## Real-world: production slog config

Tag the default logger at startup so every log line — including
third-party libraries that use `slog.Default()` — carries the build
group:

```go
import (
    "log/slog"
    "os"

    bislog "github.com/ubgo/buildinfo/contrib/buildinfo-slog"
)

func main() {
    handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level:     slog.LevelInfo,
        AddSource: false,
    })
    logger := slog.New(handler).With(bislog.Group())
    slog.SetDefault(logger)

    slog.Info("startup")
    // → {"time":"…","level":"INFO","msg":"startup","build":{"version":"1.2.3",...}}
}
```

`With` is called once at boot. The resulting logger is the new default,
so every package that calls `slog.Info` / `slog.Error` automatically
gets the build group attached.

## Combining with `-ldflags` stamping

Build with `-ldflags` populated values (see [`02-ldflags-stamping`](../02-ldflags-stamping))
and every log line becomes correlatable to a deploy:

```json
{"time":"…","level":"INFO","msg":"startup","build":{"version":"1.2.3","commit":"abc1234","branch":"main",...}}
```

This is exactly what makes "show me errors from version 1.2.3 in
Loki" a one-line LogQL query.
