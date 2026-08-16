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

# fail if the module ever grows an external dependency
deps:
    #!/usr/bin/env bash
    set -euo pipefail
    external="$(go list -deps ./... | grep -E '^[^/]+\.[^/]+/' | grep -v '^github.com/1995parham-teaching/go-lecture' || true)"
    if [ -n "$external" ]; then
      echo "this module is meant to have no external dependencies, but it imports:"
      echo "$external"
      exit 1
    fi
    echo "no external dependencies"

# everything the ci runs
check: fmt-check vet deps build test

# run a single example, e.g. just run 05-strings
run example:
    go run ./{{ example }}
