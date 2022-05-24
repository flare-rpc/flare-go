package protocol

import "sync"


var msgPool = sync.Pool{
	New: func() interface{} {
		header := Header([12]byte{})
		setMagicString(&header)
		req := &RpcRequestMeta{}
		res :=  &RpcResponseMeta{}
		return &Message{
			Header: &header,
			Meta: RpcMeta{
				Request: req,
				Response: res,
			},
		}
	},
}

// GetPooledMsg gets a pooled message.
func GetPooledMsg() *Message {
	return msgPool.Get().(*Message)
}

// FreeMsg puts a msg into the pool.
func FreeMsg(msg *Message) {
	if msg != nil && cap(msg.data) < 1024 {
		msg.Reset()
		msgPool.Put(msg)
	}
}

var poolUint32Data = sync.Pool{
	New: func() interface{} {
		data := make([]byte, 4)
		return &data
	},
}
