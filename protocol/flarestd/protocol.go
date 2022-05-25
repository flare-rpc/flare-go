package fstd

import (
	"github.com/flare-rpc/flarego"
)

const (
	ProtocolName = "flare-std"
)

type protocol struct {
}

func (p *protocol) Dial(target string, options ...flarego.DialOption) (flarego.ClientConn, error) {
	return dial(target, options...)
}

func (p *protocol) NewServer(options ...flarego.ServerOption) flarego.Server {
	return newServer()
}

func init() {
	flarego.RegisterProtocol(ProtocolName, &protocol{})
}
