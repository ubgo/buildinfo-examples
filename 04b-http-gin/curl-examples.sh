#!/usr/bin/env bash
set -euo pipefail

KEY="secret-rotate-me"

echo "== public =="
curl -sS http://localhost:8080/version | jq .version

echo
echo "== internal (no key, expect 401) =="
curl -sS -o /dev/null -w "HTTP %{http_code}\n" http://localhost:8080/internal/version

echo
echo "== internal (with key) =="
curl -sS -H "X-Internal-Key: $KEY" http://localhost:8080/internal/version | jq .version
