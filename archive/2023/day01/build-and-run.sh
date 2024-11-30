#!/bin/sh
set -e

day=$(basename "$(pwd)")
part=${1:-part1}

docker build -t "$day" .
docker run -e part="$part" "$day"