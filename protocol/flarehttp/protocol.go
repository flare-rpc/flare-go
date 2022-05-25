package bhttp

import "github.com/flare-rpc/flarego"

const (
	ProtocolName = "flare-http"
)

type protocol struct {
}

func (p *protocol) Dial(target string, options ...flarego.DialOption) (flarego.ClientConn, error) {
	return dial(target, options...)
}

func (p *protocol) NewServer(options ...flarego.ServerOption) flarego.Server {
	panic("not implemented")
}

func init() {
	flarego.RegisterProtocol(ProtocolName, &protocol{})
}
