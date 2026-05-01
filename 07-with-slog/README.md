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
