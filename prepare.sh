#!/bin/bash

[[ ! -d ./vendor ]] && sh vendor.sh
docker build -t tfc/tfc-cap-updater -f Dockerfile-build .