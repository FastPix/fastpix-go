#!/usr/bin/env bash
#
# Compile-checks every example against the SDK in this checkout.
# Each example is its own package under examples/, so a plain build covers them
# all. Exits non-zero if any example fails to compile. Run from anywhere.
set -euo pipefail

MOD_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

fail=0
for dir in "$MOD_ROOT"/examples/*/; do
  name="$(basename "$dir")"
  if out="$(cd "$MOD_ROOT" && go build -o /dev/null "./examples/$name" 2>&1)"; then
    echo "OK   $name"
  else
    echo "FAIL $name"
    echo "$out"
    fail=1
  fi
done

exit $fail
