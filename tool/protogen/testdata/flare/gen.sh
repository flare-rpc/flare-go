#!/bin/sh

protoc -I. -I${GOPATH}/src \
  --go_out=. --go_opt=paths=source_relative \
  --flarego_out=. --flarego_opt=paths=source_relative helloworld.proto
