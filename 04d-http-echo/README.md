# 04d-http-echo

Expose `/version` via Echo.

## Run

```sh
go run .
curl -sS http://localhost:8080/version | jq .
```

## Expected output

```json
{
  "version": "dev",
  "commit": "...",
  "build_time": "...",
  "branch": "unknown",
  "go_version": "go1.24.x",
  "goos": "darwin",
  "goarch": "arm64",
  "modified": false,
  "modules": [...]
}
```

## What it demonstrates

The Echo adapter's `Mount(e *echo.Echo, opts ...)` follows the same shape as every other framework adapter: `WithPath` for the route and `WithMiddleware` for `echo.MiddlewareFunc`.

## With `-ldflags` stamping

Same pattern as [`02-ldflags-stamping`](../02-ldflags-stamping):

```sh
go build -ldflags "\
  -X github.com/ubgo/buildinfo.Version=1.2.3 \
  -X github.com/ubgo/buildinfo.Commit=$(git rev-parse HEAD) \
  -X github.com/ubgo/buildinfo.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
  -X github.com/ubgo/buildinfo.Branch=$(git rev-parse --abbrev-ref HEAD)"

./04d-http-echo
curl -sS http://localhost:8080/version | jq .version
# "1.2.3"
```

## Same shape across frameworks

See [`04-http-nethttp`](../04-http-nethttp) for the canonical example with curl scripts. The other framework variants (`gin`, `chi`, `fiber`) share the same `Mount` / `WithPath` / `WithMiddleware` API.
