#!/usr/bin/env bash

GENERATED_DIR=internal/rpc/models
PROTO_DIR=internal/rpc/proto

if [ ! -d "$GENERATED_DIR" ]; then \
    mkdir -p $GENERATED_DIR; \
fi

protoc --proto_path=$PROTO_DIR \
    --go_out=$GENERATED_DIR \
    --go-grpc_out=$GENERATED_DIR \
    --grpc-gateway_out $GENERATED_DIR \
    $PROTO_DIR/*.proto