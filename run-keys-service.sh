#!/usr/bin/env bash
set -e

docker build --force-rm --no-cache -t tacc-keys .

docker run --rm -p 8000:8000 tacc-keys
