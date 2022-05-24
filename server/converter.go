package server

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/flare-rpc/flarego/protocol"
	"github.com/flare-rpc/flarego/share"
)

const (
	XVersion           = "X-FLARE-Version"
	XMessageType       = "X-FLARE-MessageType"
	XHeartbeat         = "X-FLARE-Heartbeat"
	XOneway            = "X-FLARE-Oneway"
	XMessageStatusType = "X-FLARE-MessageStatusType"
	XSerializeType     = "X-FLARE-SerializeType"
	XMessageID         = "X-FLARE-MessageID"
	XServicePath       = "X-FLARE-ServicePath"
	XServiceMethod     = "X-FLARE-ServiceMethod"
	XMeta              = "X-FLARE-Meta"
	XErrorMessage      = "X-FLARE-ErrorMessage"
)

// HTTPRequest2FlareRequest converts a http request to a flare request.
func HTTPRequest2FlareRequest(r *http.Request) (*protocol.Message, error) {
	req := protocol.GetPooledMsg()

	h := r.Header
	seq := h.Get(XMessageID)
	if seq != "" {
		id, err := strconv.ParseInt(seq, 10, 64)
		if err != nil {
			return nil, err
		}
		req.SetCorrelationId(id)
	}

	meta := h.Get(XMeta)
	if meta != "" {
		metadata, err := url.ParseQuery(meta)
		if err != nil {
			return nil, err
		}
		mm := make(map[string]string)
		for k, v := range metadata {
			if len(v) > 0 {
				mm[k] = v[0]
			}
		}
		req.Metadata = mm
	}

	auth := h.Get("Authorization")
	if auth != "" {
		if req.Metadata == nil {
			req.Metadata = make(map[string]string)
		}
		req.Metadata[share.AuthKey] = auth
	}

	req.SetServiceName(h.Get(XServicePath))

	req.SetServiceMethod(h.Get(XServiceMethod))

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	req.Payload = payload

	return req, nil
}

// func FlareResponse2HttpResponse(res *protocol.Message) (url.Values, []byte, error) {
// 	m := make(url.Values)
// 	m.Set(XVersion, strconv.Itoa(int(res.Version())))
// 	if res.IsHeartbeat() {
// 		m.Set(XHeartbeat, "true")
// 	}
// 	if res.IsOneway() {
// 		m.Set(XOneway, "true")
// 	}
// 	if res.MessageStatusType() == protocol.Error {
// 		m.Set(XMessageStatusType, "Error")
// 	} else {
// 		m.Set(XMessageStatusType, "Normal")
// 	}

// 	if res.CompressType() == protocol.Gzip {
// 		m.Set("Content-Encoding", "gzip")
// 	}

// 	m.Set(XSerializeType, strconv.Itoa(int(res.SerializeType())))
// 	m.Set(XMessageID, strconv.FormatUint(res.GetRequestId(), 10))
// 	m.Set(XServicePath, res.ServicePath)
// 	m.Set(XServiceMethod, res.ServiceMethod)

// 	return m, res.Payload, nil
// }
