#!/usr/bin/env bash

if [ "$#" -lt 1 ]; then
    echo "use: $0 with args"
    echo "$0 [module]"
    echo "module:"
    echo " - common"
    echo " - posts"
    echo " - products"
    exit 0
fi


white_list=("common" "posts" "products")
input_value="$1"

found=0
for value in "${white_list[@]}"; do
    if [ "$input_value" = "$value" ]; then
        found=1
        break
    fi
done

if [ $found -eq 0 ]; then
    echo "modules not found"
    exit 1
fi

go test -v ./testing/$1
