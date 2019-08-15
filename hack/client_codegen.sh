#!/usr/bin/env bash

set -e

DIR=`dirname $0`
ROOT_DIR=$(cd "$DIR/.." && pwd)
SWAGGER_SPEC="$ROOT_DIR/swagger.json"
BITRISE_BASEDIR="$ROOT_DIR/bitrise"

GO_PKG=github.com/charliekenney23/bitrisectl

REMOTE_SWAGGER_SPEC=https://api-docs.bitrise.io/docs/swagger.json

GO_SWAGGER_IMG=quay.io/goswagger/swagger

fetch_swagger_spec() {
  echo 'Fetching Bitrise swagger spec...'

  if [ ! -f "$SWAGGER_SPEC" ]; then
    curl -o "$SWAGGER_SPEC" "$REMOTE_SWAGGER_SPEC"
  fi
}

ensure_clean_bitrise_dir() {
  rm -rf bitrise
  mkdir bitrise
}

generate_client() {
  ensure_clean_bitrise_dir

  echo "Generating Bitrise API client in '$ROOT_DIR'..."

  local workdir="/go/src/$GO_PKG"
  docker run --rm -it \
    -e GOPATH=/go \
    -e GO111MODULE=off \
    -v "$ROOT_DIR":"$workdir" \
    -w "$workdir" \
    "$GO_SWAGGER_IMG" generate client \
    --target=bitrise
}

fetch_swagger_spec
generate_client
