#!/bin/bash

# https://hub.docker.com/_/minio

docker run --name firmware-store \
  --publish 9000:9000 \
  --publish 9001:9001 \
  --volume $(pwd)/temp/firmware-store:/data \
  bitnami/minio:2021.11.9
