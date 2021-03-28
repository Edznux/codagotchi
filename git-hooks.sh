#!/usr/bin/env bash

VERSION=$(git rev-parse --short HEAD)

docker build -t codagotchi:$VERSION .
docker stop codagotchi
docker rm -f codagotchi

docker run -d \
--net host \
-v $(pwd)/save.json:/app/save.json \
--name codagotchi \
codagotchi:$VERSION

# we "overwrite" ourselve in case the hook was updated
./install-hooks.sh 