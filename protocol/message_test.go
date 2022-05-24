package protocol

import (
	"bytes"
	"fmt"
	"github.com/flare-rpc/flarego/codec"
	"testing"
)

func TestMessage(t *testing.T) {
	req := NewMessage()
	req.SetCompressType(CompressType_COMPRESS_TYPE_NONE)
	req.SetResponseStatus(Normal)

	req.SetCorrelationId(1234567890)

	m := make(map[string]string)
	req.SetServiceName("Arith")
	req.SetServiceMethod("Add")
	m["__ID"] = "6ba7b810-9dad-11d1-80b4-00c04fd430c9"
	req.Metadata = m

	pb := TestProto{
		A: 1,
		B: 2,
	}

	coder := codec.PBCodec{}
	payload, _ :=coder.Encode(&pb)

	req.Payload = []byte(payload)

	var buf bytes.Buffer
	_, err := req.WriteTo(&buf)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("buffer len: ", buf.Len())

	res, err := Read(&buf)
	if err != nil {
		t.Fatal(err)
	}

	if req.GetCorrelationId() != 1234567890 {
		t.Errorf("expect 1234567890 but got %d", res.GetCorrelationId())
	}

	if req.GetServiceName() != "Arith" || req.GetServiceMethod() != "Add" || req.Metadata["__ID"] != "6ba7b810-9dad-11d1-80b4-00c04fd430c9" {
		t.Errorf("got wrong metadata: %v", res.Metadata)
	}

	if res.GetCorrelationId() != 1234567890 {
		t.Errorf("expect 1234567890 but got %d", res.GetCorrelationId())
	}


	if res.GetServiceName() != "Arith" || res.GetServiceMethod() != "Add" || res.Metadata["__ID"] != "6ba7b810-9dad-11d1-80b4-00c04fd430c9" {
		t.Errorf("got wrong metadata: %v", res.Metadata)
	}

	if string(res.Payload) != string(payload) {
		t.Errorf("got wrong payload: %v", string(res.Payload))
	}
}
