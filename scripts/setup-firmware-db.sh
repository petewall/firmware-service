#!/bin/bash

# https://hub.docker.com/_/mongo?tab=description

docker run \
  --name firmware-db \
  --env MONGO_INITDB_ROOT_USERNAME=mongoadmin \
  --env MONGO_INITDB_ROOT_PASSWORD=secret \
  # --volume $(pwd)/temp/firmware-db:/data/db
  mongo:5.0.4
