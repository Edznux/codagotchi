#!/usr/bin/env bash
docker build -t codagotchi:latest .
docker stop -f codagotchi
docker run -d --rm -p 8080:8080 --name codagotchi codagotchi:latest