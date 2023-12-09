#!/usr/bin/env bash

if [ "$#" -lt 2 ]; then
    echo "use: $0 with args [DOCKER_USERNAME DOCKER_IMAGE_NAME]"
    exit 1
fi

DOCKER_USERNAME=$1
DOCKER_IMAGE_NAME=$2

LATEST_TAG="$DOCKER_USERNAME/$DOCKER_IMAGE_NAME:latest"

docker rmi -f $LATEST_TAG

