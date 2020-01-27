#!/bin/bash

mkdir -p src/vendor
docker run --rm -v "$(pwd)/src":/home -w /home --entrypoint "sh" golang:1.13.6-alpine3.11 -c "go mod vendor"
