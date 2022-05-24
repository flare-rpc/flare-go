#!/bin/sh

protoc -I. -I${GOPATH}/src \
  --gofast_out=. --gofast_opt=paths=source_relative \
  --flarego_out=. --flarego_opt=paths=source_relative helloworld.proto
