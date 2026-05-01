# 03-modules-list

Dump every dependency module the binary linked, with its version and module sum. SBOM-friendly.

## What it demonstrates

- `Info.Modules` is populated automatically by `runtime/debug.ReadBuildInfo`.
- Each entry shows `Path`, `Version`, and `Sum` (the `h1:...` hash from `go.sum`).
- Modules replaced via `go.mod` `replace` directives are surfaced with the **replacement** values (not the original), so the dump reflects what was actually linked.

## Run

```sh
go run .
```

## Expected output

```
--- N dependency modules ---

PATH                                                    VERSION  SUM
github.com/ubgo/buildinfo                               v0.0.0
...
```

(For this minimal example, the only dependency is `github.com/ubgo/buildinfo` itself, but a real service would list every transitive dependency. Try it on a larger codebase.)

## Use case: lightweight SBOM endpoint

Combine with [`04-http-nethttp`](../04-http-nethttp) or any other adapter to expose the modules list at a private `/internal/sbom` endpoint behind auth middleware. Useful for compliance reporting without pulling in a full SBOM toolchain.

```go
mux.Handle("/internal/sbom", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
    json.NewEncoder(w).Encode(buildinfo.Get().Modules)
})))
```
