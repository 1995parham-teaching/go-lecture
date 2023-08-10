default:
    @just --list

# build steps
build:
    #!/usr/bin/env bash
    set -euxo pipefail
    for step in $(find . -maxdepth 1 -type d -name '[!.]*'); do
      cd $step
      go build -o main.out || echo "$step: failed to compile"
      cd -
    done
