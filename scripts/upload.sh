#!/usr/bin/env bash

if [ -z "${3}" ]; then
  echo "USAGE: upload.sh <type> <version> <path/to/firmware>" >&2
  exit 1
fi

type=$1
version=$2
firmware=$3

if [ ! -f "${firmware}" ]; then
  echo "${firmware} is not a file" >&2
  exit 1
fi

curl \
  --verbose \
  --fail \
  --upload-file "${firmware}" \
  --header "Content-Type: application/octet-stream" \
  "${FIRMWARE_SERVICE:-localhost:5000}/api/firmware/${type}/${version}"
