#!/bin/bash

[[ ! -d src/vendor ]] && sh vendor.sh
IMAGE=tfc/cap-updater
VERSION=

docker build -t ${IMAGE} -f Dockerfile-build .
docker tag ${IMAGE} localhost:5000/${IMAGE}${VERSION}
docker push localhost:5000/${IMAGE}${VERSION}
