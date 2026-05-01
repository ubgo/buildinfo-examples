# 01-basic

The simplest possible use of `ubgo/buildinfo` — print `buildinfo.Get()` to stdout.

## What it demonstrates

- Zero-config: no `-ldflags`, no setup.
- Auto-population from `runtime/debug.ReadBuildInfo` — `Commit`, `BuildTime`, `Modified` are filled from VCS metadata Go embeds since 1.18.
- Sentinel defaults: `Version` falls back to `"dev"`, `Branch` to `"unknown"` when nothing else fills them.

## Run

```sh
go run .
```

## Expected output

```
--- flat fields ---
version:     dev
commit:      <git-revision-hash>
build_time:  <vcs-time>
branch:      unknown
go_version:  go1.24.x
goos/goarch: <your os>/<your arch>
modified:    <true if your tree is dirty>
modules:     <N> entries

--- JSON ---
{
  "version": "dev",
  "commit": "...",
  "build_time": "...",
  "branch": "unknown",
  "go_version": "go1.24.x",
  ...
}
```

`Branch` is `"unknown"` because the toolchain does not embed branch information. To populate it, see [`02-ldflags-stamping`](../02-ldflags-stamping).
