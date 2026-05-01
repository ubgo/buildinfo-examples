# 04c-http-chi

Expose `/version` via Chi.

## What it demonstrates

- `bchi.Mount(r)` registers the version handler on any `chi.Router`.
- Chi's stdlib-compatible `func(http.Handler) http.Handler` middleware shape.
- Composing with `middleware.Recoverer` from chi's stock middleware.

## Run

```sh
go run .
curl http://localhost:8080/version | jq .
```
