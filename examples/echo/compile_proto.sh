#!/bin/bash

protoc --go_out=. --go_opt=paths=source_relative --flarego_out=. --flarego_opt=paths=source_relative  *.proto