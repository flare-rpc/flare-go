#!/bin/sh

protoc -I. -I${GOPATH}/src \
  --go_out=. --go_opt=paths=source_relative \
  --flare_out=. --flare_opt=paths=source_relative helloworld.proto
