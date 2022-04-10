#!/bin/bash

# See the following for an example of these commands: https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code

# Using deprecated gRPC plugin
#protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

# Using new gRPC plugin
protoc --go_out=. --go-grpc_out=. greet/greetpb/greet.proto
protoc --go_out=. --go-grpc_out=. calculator/calculatorpb/calculator.proto
