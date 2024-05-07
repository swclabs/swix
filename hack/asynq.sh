#!/usr/bin/env bash

if [ "$#" -lt 1 ]; then
    echo "use: $0 with args"
    echo "hack/asynq [uri] [password]"
    exit 0
fi

asynq dash --uri $1 --password $2
