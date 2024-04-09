default:
    @just --list

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
