package main

import (
	"context"
	"fmt"
	"net"

	"github.com/flare-rpc/flarego"
	"github.com/flare-rpc/flarego/examples/echo"
	"github.com/flare-rpc/flarego/protocol/flarestd"
)

type echoService struct {
}

func (s *echoService) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	fmt.Println("client call")
	return &echo.EchoResponse{
		Message: "reply: " + req.Message,
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := flarego.NewServer(flarestd.ProtocolName)
	echo.RegisterEchoServerServer(server, &echoService{})
	server.Serve(l)
}
