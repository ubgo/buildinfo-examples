# 04e-http-fiber

Expose `/version` via Fiber.

## Run

```sh
go run .
curl http://localhost:8080/version | jq .
```

The Fiber adapter's `Mount(r fiber.Router, opts ...)` accepts any `fiber.Router`, so mounting on a sub-group (`app.Group("/api/v1")`) works the same.
