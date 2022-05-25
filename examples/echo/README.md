# Usage

- Install the protobuf compiler https://grpc.io/docs/protoc-installation/
- Install the Go protobuf plugin and flare plugin

``` bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install github.com/flare-rpc/flarego/protoc-gen-flarego@latest
```

- Generate protobuf files

``` bash
protoc --go_out=. --go_opt=paths=source_relative --flarego_out=. --flarego_opt=paths=source_relative  *.proto
```

For how to build server and client code, see [server.go](./server/server.go) and [client.go](./client/client.go)