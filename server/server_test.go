package server

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"testing"
	"time"

	testutils "github.com/flare-rpc/flarego/_testutils"
	"github.com/flare-rpc/flarego/codec"
	"github.com/flare-rpc/flarego/protocol"
	"github.com/flare-rpc/flarego/share"
	"github.com/stretchr/testify/assert"
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Println("server Mul")
	return nil
}

func (t *Arith) ThriftMul(ctx context.Context, args *testutils.ThriftArgs_, reply *testutils.ThriftReply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) ConsumingOperation(ctx context.Context, args *testutils.ThriftArgs_, reply *testutils.ThriftReply) error {
	reply.C = args.A * args.B
	time.Sleep(10 * time.Second)
	return nil
}

func TestShutdownHook(t *testing.T) {
	s := NewServer()
	var cancel1 context.CancelFunc
	s.RegisterOnShutdown(func(s *Server) {
		var ctx context.Context
		ctx, cancel1 = context.WithTimeout(context.Background(), 155*time.Second)
		s.Shutdown(ctx)
	})
	s.RegisterName("Arith2", new(Arith), "")
	s.Register(new(Arith), "")
	go s.Serve("tcp", ":0")

	time.Sleep(time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	s.Shutdown(ctx)
	cancel()
	if cancel1 != nil {
		cancel1()
	}
}

func TestHandler(t *testing.T) {
	// use jsoncodec

	req := protocol.NewMessage()
	req.SetCompressType(protocol.CompressType_COMPRESS_TYPE_NONE)
	req.SetResponseStatus(protocol.Normal)
	req.SetCorrelationId(1234567890)

	req.SetServiceName("Arith")
	req.SetServiceMethod("Mul")

	argv := &Args{
		A: 10,
		B: 20,
	}
	coder := codec.PBCodec{}
	data, err := coder.Encode(argv)
	if err != nil {
		t.Fatal(err)
	}

	req.Payload = data

	serverConn, clientConn := net.Pipe()

	handler := func(ctx *Context) error {
		req := &Args{}
		res := &Reply{}
		ctx.Bind(req)
		res.C = req.A * req.B

		return ctx.Write(res)
	}

	go func() {
		ctx := NewContext(share.NewContext(context.Background()), serverConn, req, nil)
		err = handler(ctx)
		assert.NoError(t, err)

		serverConn.Close()
	}()

	data, err = ioutil.ReadAll(clientConn)
	assert.NoError(t, err)

	resp, err := protocol.Read(bytes.NewReader(data))
	assert.NoError(t, err)

	assert.Equal(t, "Arith", resp.GetServiceName())
	assert.Equal(t, "Mul", resp.GetServiceMethod())
	assert.Equal(t, req.GetCorrelationId(), resp.GetCorrelationId())

	r := Reply{C:200}
	rd, _:=coder.Encode(&r)
	assert.Equal(t, string(rd), string(resp.Payload))
}

func Test_validIP6(t *testing.T) {
	type args struct {
		ipAddress string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{ipAddress: "[CDCD:910A:2222:5498:8475:1111:3900:2020]:8080"}, true},
		{"2", args{ipAddress: "[1030::C9B4:FF12:48AA:1A2B]:8080"}, true},
		{"3", args{ipAddress: "[2000:0:0:0:0:0:0:1]:8080"}, true},
		{"4", args{ipAddress: "127.0.0.1:8080"}, false},
		{"5", args{ipAddress: "localhost:8080"}, false},
		{"6", args{ipAddress: "127.1:8080"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, validIP6(tt.args.ipAddress), "validIP6(%v)", tt.args.ipAddress)
		})
	}
}
