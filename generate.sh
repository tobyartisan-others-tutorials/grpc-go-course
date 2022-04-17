#!/bin/bash

# See the following for an example of these commands: https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code

# Using new gRPC plugin
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/tobyartisan-others-tutorials/grpc-go-course  --go-grpc_out=. --go-grpc_opt=module=github.com/tobyartisan-others-tutorials/grpc-go-course greet/proto/dummy.proto
