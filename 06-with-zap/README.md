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
