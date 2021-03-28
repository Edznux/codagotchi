#!/usr/bin/env bash


docker build -t codagotchi:latest .
docker stop codagotchi
docker rm -f codagotchi
docker run -d \
-p 8080:8080 \
-v $(pwd)/save.json:/app/save.json \
--name codagotchi \
codagotchi:latest

# we "overwrite" ourselve in case the hook was updated
./install-hooks.sh 