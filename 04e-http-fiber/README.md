# 04e-http-fiber

Expose `/version` via Fiber.

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

The Fiber adapter's `Mount(r fiber.Router, opts ...)` accepts any `fiber.Router`, so mounting on a sub-group (`app.Group("/api/v1")`) works the same. Middleware is `fiber.Handler`, applied via `WithMiddleware` in declaration order.

## With `-ldflags` stamping

Same pattern as [`02-ldflags-stamping`](../02-ldflags-stamping):

```sh
go build -ldflags "\
  -X github.com/ubgo/buildinfo.Version=1.2.3 \
  -X github.com/ubgo/buildinfo.Commit=$(git rev-parse HEAD) \
  -X github.com/ubgo/buildinfo.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
  -X github.com/ubgo/buildinfo.Branch=$(git rev-parse --abbrev-ref HEAD)"

./04e-http-fiber
curl -sS http://localhost:8080/version | jq .version
# "1.2.3"
```

## Same shape across frameworks

See [`04-http-nethttp`](../04-http-nethttp) for the canonical example with curl scripts. The other framework variants (`gin`, `chi`, `echo`) share the same `Mount` / `WithPath` / `WithMiddleware` API.
