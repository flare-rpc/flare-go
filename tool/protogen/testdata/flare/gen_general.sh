#!/bin/sh

protoc -I. -I${GOPATH}/src \
  --gofast_out=. --gofast_opt=paths=source_relative \
  --flare_out=. --flare_opt=paths=source_relative helloworld.proto
