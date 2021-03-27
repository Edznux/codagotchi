#!/usr/bin/env bash
docker build -t codagotchi:latest .
docker stop -f codagotchi
docker run -d -p 8080:8080 -v $(pwd)/save.json:/app/save.json --name codagotchi codagotchi:latest