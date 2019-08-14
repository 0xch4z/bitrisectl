#!/usr/bin/env bash

set -e

DIR=`dirname $0`
ROOT_DIR=$(cd "$DIR/.." && pwd)
SWAGGER_SPEC="$ROOT_DIR/swagger.json"
BITRISE_BASEDIR="$ROOT_DIR/bitrise"

echo $DIR $ROOT_DIR $SWAGGER_SPEC

REMOTE_SWAGGER_SPEC='https://api-docs.bitrise.io/docs/swagger.json'

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

  docker run --rm -it \
    -e GO111MODULE=on \
    -v "$ROOT_DIR":/var/work \
    -w /var/work \
    "$GO_SWAGGER_IMG" generate client \
    --target=/var/work/bitrise
}

fetch_swagger_spec
generate_client
