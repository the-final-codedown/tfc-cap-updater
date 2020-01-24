#!/bin/bash

docker run --rm --name tfc-cap-updater -p 50051:50051 -e DB_HOST=mongo --link mongo "$@" tfc/tfc-cap-updater