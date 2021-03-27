#!/usr/bin/env bash
docker build -t edznux/codagotchi .
docker stop codagotchi
docker rmi edznux/codagotchi
docker run -d --rm -p 8080:8080 --name codagotchi edznux/codagotchi