# Contributing to ubgo/buildinfo-examples

Thanks for your interest. This repository is licensed under the **Apache License 2.0**. Pull requests are welcome.

## When to open an issue first

Open a GitHub issue **before** sending a non-trivial PR if you want to:

- Add a new example that demonstrates a feature not currently covered.
- Refactor multiple examples in one go.
- Change the per-example file conventions.

For tiny fixes (typos, broken commands, version bumps), feel free to open a PR directly.

## Adding a new example

Each example is a self-contained Go module under its own numbered subdirectory.

1. Pick the next available number that fits the learning path (see [`README.md`](./README.md)).
2. Create the directory: `NN-short-name/`.
3. Add the standard files:
   - `go.mod` — pins specific upstream versions, plus a `replace` directive pointing at a sibling local clone of `ubgo/buildinfo` if working from source.
   - `main.go` — single-file runnable. Keep it under 100 lines.
   - `README.md` — what it demonstrates, prerequisites, run command, expected output.
   - `expected_output.txt` — optional; when present, CI uses it to smoke-test the run.
   - `Makefile` / `docker-compose.yml` — only when the example genuinely needs them.
4. Add the example to the `build` and (if applicable) `smoke-test` matrices in `.github/workflows/test.yml`.
5. Run `task ci` locally before pushing.

## Conventions

- **Single-file `main.go`.** Examples are about clarity, not modularity. If you find yourself wanting an internal package, the example is too big.
- **Minimal dependencies.** Each example pulls only `ubgo/buildinfo`, the relevant adapter, and the framework being demonstrated. Avoid testing libraries, mock frameworks, or extra observability stacks unless that's literally what the example demonstrates.
- **Pin exact versions** in `go.mod`. No floating ranges. Examples drift over time — pinning makes drift detectable.
- **Conventional Commits** for PR titles: `feat(examples): add 08-with-pyroscope`, `fix(04-http-nethttp): handle missing /version path`, `docs(readme): clarify ldflags syntax`.

## License of contributions

By submitting a pull request, you agree that your contribution is provided under the same Apache License 2.0 as the rest of the repository (per the standard "inbound = outbound" rule, codified in section 5 of the Apache License).
