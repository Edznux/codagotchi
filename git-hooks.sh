#!/usr/bin/env bash


docker build -t codagotchi:latest .
docker stop codagotchi
docker rm -f codagotchi
docker run -d \
--net host \
-v $(pwd)/save.json:/app/save.json \
--name codagotchi \
codagotchi:latest

# we "overwrite" ourselve in case the hook was updated
./install-hooks.sh 