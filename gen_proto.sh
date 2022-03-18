#! /bin/bash
set -e
echo ">>> Generate authgrpc Proto."
protoc -I ./authgrpc --go_out ./authgrpc --go_opt paths=source_relative \
    --go-grpc_out=./authgrpc --go-grpc_opt paths=source_relative \
    ./authgrpc/authgrpc.proto
echo ">>> Generate authgrpc Proto OK."