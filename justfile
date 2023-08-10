default:
    @just --list

# build steps
build:
    #!/usr/bin/env bash
    set -euo pipefail
    for step in $(find . -maxdepth 1 -type d -name '[!.]*'); do
      echo "*** $step ***"
      echo "*************"
      cd $step &> /dev/null
      go build -o main.out || true
      cd - &> /dev/null
    done
