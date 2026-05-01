# 02-ldflags-stamping

CI-style build-time version stamping. Demonstrates how `-ldflags` overrides win over `runtime/debug` VCS data and how to wire them through a `Makefile`.

## What it demonstrates

- The four `-ldflags` overrides: `Version`, `Commit`, `BuildTime`, `Branch`.
- Demo values supplied via Makefile variables (`VERSION`, `COMMIT`, `BUILD_TIME`, `BRANCH`).
- A `build-from-git` target that pulls real values from `git describe`, `git rev-parse`, and `date`.

## Run

```sh
make build           # builds with demo values: 1.2.3 / abc123def456 / ...
make run             # builds + runs
```

```sh
$ ./app
version=1.2.3 commit=abc123def456 build_time=2026-04-26T12:00:00Z branch=main
```

Override values from the command line:

```sh
make build VERSION=2.0.0-beta.1 BRANCH=release/2.0
./app
# version=2.0.0-beta.1 commit=abc123def456 build_time=2026-04-26T12:00:00Z branch=release/2.0
```

Build with real git metadata:

```sh
make build-from-git
./app
# version=v0.1.0 commit=<hash> build_time=<utc-now> branch=<your-branch>
```

## How it works

The Go toolchain's `-ldflags="-X path.var=value"` overwrites the named `var`-level string variable at link time. `buildinfo` declares those variables in `ldflags.go`:

```go
package buildinfo
var (
    Version   string
    Commit    string
    BuildTime string
    Branch    string
)
```

When `Get()` is called, it reads `runtime/debug.ReadBuildInfo` first to populate VCS data, then overlays the `-ldflags` variables on top. CI-stamped values always win.
