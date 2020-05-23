#!/bin/sh

echo "Generating root gRPC server protos"

# Generate the protos.
protoc -I/usr/local/include -I. \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --go_out=plugins=grpc,paths=source_relative:. \
       accounts.proto
