package serverplugin

import (
	"testing"
	"time"

	"github.com/flare-rpc/flarego/server"
	metrics "github.com/rcrowley/go-metrics"
)

func TestConsulRegistry(t *testing.T) {
	s := server.NewServer()

	r := &ConsulRegisterPlugin{
		ServiceAddress: "tcp@127.0.0.1:8972",
		ConsulServers:  []string{"127.0.0.1:8500"},
		BasePath:       "/flare_test",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		return
	}
	s.Plugins.Add(r)

	s.RegisterName("Arith", new(Arith), "")
	go s.Serve("tcp", "127.0.0.1:8972")
	defer s.Close()

	if len(r.Services) != 1 {
		t.Fatal("failed to register services in consul")
	}

	if err := r.Stop(); err != nil {
		t.Fatal(err)
	}
}
