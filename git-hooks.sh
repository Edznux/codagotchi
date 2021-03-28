#!/usr/bin/env bash


docker build -t codagotchi:latest .
docker stop codagotchi
docker rm codagotchi
docker run -d \
-p 8080:8080 \
-v /tmp/.X11-unix:/tmp/.X11-unix \
-e DISPLAY=${DISPLAY} \
-v $(pwd)/save.json:/app/save.json \
--name codagotchi \
codagotchi:latest

# we "overwrite" ourselve in case the hook was updated
./install-hooks.sh 