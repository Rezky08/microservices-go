#!/bin/bash
protoc --proto_path=protos protos/*.proto --go_out=. --go-grpc_out=.
#protoc "--proto_path=cmd/services/$1/protos/" --proto_path=protos "cmd/services/$1/protos/*.proto" "--go_out=cmd/services/$1" "--go-grpc_out=cmd/services/$1"