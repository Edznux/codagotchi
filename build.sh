#!/usr/bin/env bash

# Build the wasm version first
GOOS=js GOARCH=wasm go build -o webserver/static/js/codagotchi.wasm .

# then build the "server" version so we can run the webserver and all
go build
