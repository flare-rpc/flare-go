#!/bin/sh

go get github.com/gogo/protobuf/protoc-gen-gogofaster
protoc -I. -I${GOPATH}/src --gogofaster_out=. --gogofaster_opt=paths=source_relative  options.proto
protoc -I. -I${GOPATH}/src --gogofaster_out=. --gogofaster_opt=paths=source_relative  streaming_rpc_meta.proto
protoc -I. -I${GOPATH}/src --gogofaster_out=. --gogofaster_opt=paths=source_relative  flare_rpc_meta.proto
protoc -I. -I${GOPATH}/src --gogofaster_out=. --gogofaster_opt=paths=source_relative   message_test.proto
