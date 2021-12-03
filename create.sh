#!/usr/bin/env bash

set -Eeuf -o pipefail
set -x

main() {
  for day in {01..25}; do
    path="d${day}"
    [ -d "${path}" ] && continue
    mkdir -p "${path}"
    cp ./d01.go "${path}/${path}.go"
    cp ./d01_test.go "${path}/${path}_test.go"
    sed -i "1 s/^\(package d\)01\$/\1${day}/" "${path}/${path}.go"
    sed -i "1 s/^\(package d\)01\$/\1${day}/" "${path}/${path}_test.go"
  done
}
main "$@"
