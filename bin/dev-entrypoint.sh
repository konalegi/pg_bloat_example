#!/bin/sh

set -e

if [ "$1" = "app" ]; then
  echo 'building app...'
  go build -o $GOPATH/bin/app
  exec app
else
  exec "$@"
fi
