#!/usr/bin/env bash
GENERATED_DIR=src
PROTO_DIR=proto

python -m grpc_tools.protoc -I. \
    --python_out=./$GENERATED_DIR \
    --grpc_python_out=./$GENERATED_DIR \
    --pyi_out=./$GENERATED_DIR \
    $PROTO_DIR/*.proto
