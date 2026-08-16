default:
    @just --list

# run the Go 1.26 modernizer fixers across the whole module
modernize:
    go fix ./...

# format every package
fmt:
    go fmt ./...

# fail when something is not gofmt clean, without changing it
fmt-check:
    #!/usr/bin/env bash
    set -euo pipefail
    unformatted="$(gofmt -l .)"
    if [ -n "$unformatted" ]; then
      echo "not gofmt clean:"
      echo "$unformatted"
      exit 1
    fi

# build every example
build:
    go build ./...

# report suspicious constructs
vet:
    go vet ./...

# run the tests
test:
    go test ./...

# everything the ci runs
check: fmt-check vet build test

# run a single example, e.g. just run 05-strings
run example:
    go run ./{{ example }}
