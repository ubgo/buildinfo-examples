# 04-http-nethttp

Expose `/version` via stdlib `net/http`. Zero third-party dependencies.

## What it demonstrates

- `binethttp.Mount(mux)` registers `GET /version` on a stdlib `*http.ServeMux`.
- Default route is `/version`. Override with `binethttp.WithPath("/...")`.
- Graceful shutdown on SIGINT / SIGTERM with a 5s deadline.

## Run

```sh
go run .
```

In another terminal:

```sh
curl -sS http://localhost:8080/version | jq .
# or use the included script:
./curl-examples.sh
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

## With `-ldflags` stamping

Combine with the technique from [`02-ldflags-stamping`](../02-ldflags-stamping):

```sh
go build -ldflags "\
  -X github.com/ubgo/buildinfo.Version=1.2.3 \
  -X github.com/ubgo/buildinfo.Commit=$(git rev-parse HEAD) \
  -X github.com/ubgo/buildinfo.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
  -X github.com/ubgo/buildinfo.Branch=$(git rev-parse --abbrev-ref HEAD)"

./04-http-nethttp
curl -sS http://localhost:8080/version | jq .version
# "1.2.3"
```
