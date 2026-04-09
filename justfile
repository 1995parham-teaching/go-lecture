default:
    @just --list

# run the Go 1.26 modernizer fixers across the whole module
modernize:
    go fix ./...

# build all the examples
build:
    #!/usr/bin/env bash
    set -euo pipefail
    for step in $(find . -maxdepth 1 -type d -name '[!.]*'); do
      echo "*** $step ***"
      echo "*************"
      cd $step &> /dev/null
      go build -o main.out || true
      go test
      cd - &> /dev/null
    done
