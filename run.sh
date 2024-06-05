#!/bin/bash
set -e

# Create a temporary file to hold the compiled binary
tmpFile=$(mktemp)

# Change to the directory of the script and build the Go project
( cd "$(dirname "$0")" && go build -o "$tmpFile" ./main.go )

# Execute the compiled binary with any passed arguments
exec "$tmpFile" "$@"
