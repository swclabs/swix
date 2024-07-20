#!/usr/bin/env bash

if [ "$#" -lt 2 ]; then
    echo "use: $0 with args [DOCKER_USERNAME DOCKER_IMAGE_NAME MODULE_NAME]"
    exit 1
fi

DOCKER_USERNAME=$1
DOCKER_IMAGE_NAME=$2
MODULE_NAME=$3

if [ "$3" == "" ]; then
    MODULE_NAME=""
    elif [ "$3" == "account-management" ]; then
    MODULE_NAME="module/accountmanagement"
    elif [ "$3" == "product-management" ]; then
    MODULE_NAME="module/productmanagement"
fi

COMMIT_ID=$(git rev-parse --short HEAD)

COMMIT_TAG="$DOCKER_USERNAME/$DOCKER_IMAGE_NAME:$COMMIT_ID"
LATEST_TAG="$DOCKER_USERNAME/$DOCKER_IMAGE_NAME:latest"

echo "build: $COMMIT_TAG"
docker build --build-arg APP_MODULE=$MODULE_NAME -t $COMMIT_TAG Dockerfile.prod