#!/bin/sh

#go get github.com/gogo/protobuf/protoc-gen-gogofaster
protoc -I. -I${GOPATH}/src --gogofaster_out=. --gogofaster_opt=paths=source_relative  args.proto

