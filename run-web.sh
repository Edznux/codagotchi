#!/usr/bin/env bash

# need to create a fake display (:99)
Xvfb :99 -screen 0 1080x720x24 &
/app/codagotchi web