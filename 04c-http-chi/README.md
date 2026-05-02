# 04c-http-chi

Expose `/version` via Chi.

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

- `bchi.Mount(r)` registers the version handler on any `chi.Router`.
- Chi's stdlib-compatible `func(http.Handler) http.Handler` middleware shape works with `bchi.WithMiddleware`.
- Composing with `chi/middleware.Recoverer` from chi's stock middleware.

## With `-ldflags` stamping

Same pattern as [`02-ldflags-stamping`](../02-ldflags-stamping):

```sh
go build -ldflags "\
  -X github.com/ubgo/buildinfo.Version=1.2.3 \
  -X github.com/ubgo/buildinfo.Commit=$(git rev-parse HEAD) \
  -X github.com/ubgo/buildinfo.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
  -X github.com/ubgo/buildinfo.Branch=$(git rev-parse --abbrev-ref HEAD)"

./04c-http-chi
curl -sS http://localhost:8080/version | jq .version
# "1.2.3"
```

## Same shape across frameworks

The contrib API (`Mount`, `WithPath`, `WithMiddleware`) is identical across `nethttp`, `gin`, `chi`, `echo`, and `fiber` — only the middleware signature differs. See [`04-http-nethttp`](../04-http-nethttp) for the canonical example with curl scripts.
