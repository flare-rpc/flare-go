package main

import (
	"context"
	"fmt"

	"github.com/flare-rpc/flarego"
	"github.com/flare-rpc/flarego/examples/echo"
	"github.com/flare-rpc/flarego/protocol/flarestd"
)

func main() {
	clientConn, err := flarego.Dial(flarestd.ProtocolName, "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	client := echo.NewEchoServerClient(clientConn)
	resp, err := client.Echo(context.Background(), &echo.EchoRequest{
		Message: "hello",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}
