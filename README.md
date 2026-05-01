# buildinfo-examples

Runnable example applications demonstrating every adapter and usage pattern of [`github.com/ubgo/buildinfo`](https://github.com/ubgo/buildinfo).

Each example is a self-contained Go module under its own subdirectory. Clone the repo, `cd` into any example, and run `go run .` to see the feature working end-to-end.

## Quick start

```sh
git clone https://github.com/ubgo/buildinfo-examples.git
cd buildinfo-examples/01-basic
go run .
```

## Learning path

### Basics

- [`01-basic`](./01-basic) — auto-populated `Info` via `runtime/debug.ReadBuildInfo`.
- [`02-ldflags-stamping`](./02-ldflags-stamping) — CI-style version stamping with `-ldflags` (Makefile included).
- [`03-modules-list`](./03-modules-list) — dump dependency modules — SBOM-friendly.

### HTTP framework adapters

Each example boots a tiny server and exposes `/version`. The shape of the API (`Handler`, `Mount`, `WithPath`, `WithMiddleware`) is uniform across all five frameworks.

- [`04-http-nethttp`](./04-http-nethttp) — stdlib `net/http`.
- [`04b-http-gin`](./04b-http-gin) — Gin + auth middleware example.
- [`04c-http-chi`](./04c-http-chi) — Chi.
- [`04d-http-echo`](./04d-http-echo) — Echo.
- [`04e-http-fiber`](./04e-http-fiber) — Fiber.

### Observability + logger adapters

- [`05-with-otel-resource`](./05-with-otel-resource) — OpenTelemetry resource attributes.
- [`06-with-zap`](./06-with-zap) — Zap log fields and namespace.
- [`07-with-slog`](./07-with-slog) — stdlib `log/slog` Attrs and group.

## Versioning

Each example's `go.mod` pins specific upstream versions of `ubgo/buildinfo` and any contrib adapters it imports. When a new minor release of the lib lands, all examples are updated in one PR.

During pre-v1.0 development, examples use `replace` directives that point at sibling local clones of `ubgo/buildinfo`. After tagging, the `replace` directives are removed.

## Compatibility

Requires Go 1.24 or later.

## License

Apache License 2.0. See [`LICENSE`](./LICENSE).
