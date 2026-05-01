# 04d-http-echo

Expose `/version` via Echo.

## Run

```sh
go run .
curl http://localhost:8080/version | jq .
```

The Echo adapter's `Mount(e *echo.Echo, opts ...)` follows the same shape as every other framework adapter: `WithPath` for the route and `WithMiddleware` for `echo.MiddlewareFunc`.
