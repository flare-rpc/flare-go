# curl -O https://raw.githubusercontent.com/rpcxio/flare-benchmark/master/proto/benchmark.proto

# generate .go files from IDL
protoc --go_out=./ ./benchmark.proto

