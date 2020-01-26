#!/bin/bash

mkdir -p vendor
docker run --rm -v "$(pwd)":/home -w /home --entrypoint "sh" golang:1.13.6-alpine3.11 -c "go mod vendor"
