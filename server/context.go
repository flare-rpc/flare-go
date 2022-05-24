package server

import (
	"github.com/flare-rpc/flarego/codec"
	"net"

	"github.com/flare-rpc/flarego/protocol"
	"github.com/flare-rpc/flarego/share"
)

// Context represents a flare FastCall context.
type Context struct {
	conn net.Conn
	req  *protocol.Message
	ctx  *share.Context

	writeCh chan *[]byte
}

// NewContext creates a server.Context for Handler.
func NewContext(ctx *share.Context, conn net.Conn, req *protocol.Message, writeCh chan *[]byte) *Context {
	return &Context{conn: conn, req: req, ctx: ctx, writeCh: writeCh}
}

// Get returns value for key.
func (ctx *Context) Get(key interface{}) interface{} {
	return ctx.ctx.Value(key)
}

// SetValue sets the kv pair.
func (ctx *Context) SetValue(key, val interface{}) {
	if key == nil || val == nil {
		return
	}
	ctx.ctx.SetValue(key, val)
}

// DeleteKey delete the kv pair by key.
func (ctx *Context) DeleteKey(key interface{}) {
	if ctx.ctx == nil || key == nil {
		return
	}
	ctx.ctx.DeleteKey(key)
}

// Payload returns the  payload.
func (ctx *Context) Payload() []byte {
	return ctx.req.Payload
}

// Metadata returns the metadata.
func (ctx *Context) Metadata() map[string]string {
	return ctx.req.Metadata
}

// ServicePath returns the ServicePath.
func (ctx *Context) ServicePath() string {
	return ctx.req.GetServiceName()
}

// ServiceMethod returns the ServiceMethod.
func (ctx *Context) ServiceMethod() string {
	return ctx.req.GetServiceMethod()
}

// Bind parses the body data and stores the result to v.
func (ctx *Context) Bind(v interface{}) error {
	req := ctx.req
	if v != nil {
		coder := codec.PBCodec{}
		err := coder.Decode(req.Payload, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ctx *Context) Write(v interface{}) error {
	req := ctx.req

	coder := codec.PBCodec{}
	res := req.Clone()

	if v != nil {
		data, err := coder.Encode(v)
		if err != nil {
			return err
		}
		res.Payload = data
	}

	resMetadata := ctx.Get(share.ResMetaDataKey)
	if resMetadata != nil {
		resMetaInCtx := resMetadata.(map[string]string)
		meta := res.Metadata
		if meta == nil {
			res.Metadata = resMetaInCtx
		} else {
			for k, v := range resMetaInCtx {
				if meta[k] == "" {
					meta[k] = v
				}
			}
		}
	}

	if len(res.Payload) > 1024 && req.GetCompressType() != protocol.CompressType_COMPRESS_TYPE_NONE {
		res.SetCompressType(req.GetCompressType())
	}
	respData := res.EncodeSlicePointer()

	var err error
	if ctx.writeCh != nil {
		ctx.writeCh <- respData
	} else {
		_, err = ctx.conn.Write(*respData)
		protocol.PutData(respData)
	}

	return err
}

func (ctx *Context) WriteError(err error) error {
	req := ctx.req
	res := req.Clone()
	resMetadata := ctx.Get(share.ResMetaDataKey)
	if resMetadata != nil {
		resMetaInCtx := resMetadata.(map[string]string)
		meta := res.Metadata
		if meta == nil {
			res.Metadata = resMetaInCtx
		} else {
			for k, v := range resMetaInCtx {
				if meta[k] == "" {
					meta[k] = v
				}
			}
		}
	}

	res.SetResponseStatus(protocol.Error)
	res.SetResponseErrorMsg(err.Error())

	respData := res.EncodeSlicePointer()
	ctx.conn.Write(*respData)
	protocol.PutData(respData)

	return nil
}
