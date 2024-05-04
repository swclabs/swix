#!/usr/bin/env bash

if [ "$#" -lt 1 ]; then
    echo "use: $0 with args [module]"
    echo "module: common"
    echo "        account"
    exit 1
fi

if [ "$1" == "common" ]; then
    go test -v ./testing/common
fi
