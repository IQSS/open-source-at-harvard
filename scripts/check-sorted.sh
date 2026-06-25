#!/bin/sh
# Usage: check-sorted.sh [--fix]
# Checks that github.tsv is sorted (case-insensitive, header preserved).
# With --fix, sorts it in place.
file=github.tsv
sorted=$({ head -1 "$file"; tail -n +2 "$file" | sort -f; })
if [ "$sorted" = "$(cat "$file")" ]; then
  exit 0
fi
if [ "$1" = "--fix" ]; then
  printf '%s\n' "$sorted" > "$file"
  echo "Sorted $file"
else
  echo "$file is not sorted. Run: scripts/check-sorted.sh --fix"
  exit 1
fi
