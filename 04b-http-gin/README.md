# 04b-http-gin

Expose `/version` via Gin, with a public route and an internal-key-protected variant.

## What it demonstrates

- `bgin.Mount(r)` registers `GET /version` on a Gin engine.
- `bgin.Mount(group, bgin.WithPath("/version"))` mounts the same handler on a group, inheriting the group's middleware (`internalKeyAuth` here).
- Same `WithPath` / `WithMiddleware` API shape as every other framework adapter.

## Run

```sh
go run .
```

In another terminal:

```sh
./curl-examples.sh
```

Or by hand:

```sh
# public — open
curl http://localhost:8080/version | jq .

# internal — needs key
curl -H "X-Internal-Key: secret-rotate-me" http://localhost:8080/internal/version | jq .

# internal — without key returns 401
curl -i http://localhost:8080/internal/version
# HTTP/1.1 401 Unauthorized
```

## Notes

- The example uses `subtle.ConstantTimeCompare` to avoid timing-leak attacks on the key — recommended for any HMAC-like comparison.
- `gin.SetMode(gin.ReleaseMode)` keeps the example output uncluttered.
