#!/usr/bin/env bash
set -euo pipefail

# Pretty-prints the /version endpoint. Requires the example to be running:
#   go run .

curl -sS http://localhost:8080/version | jq .
