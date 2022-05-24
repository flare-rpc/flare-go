package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/flare-rpc/flarego/codec"
	"github.com/flare-rpc/flarego/util"
	"github.com/valyala/bytebufferpool"
	"io"
)

const (
	// Normal is normal requests and responses.
	Normal int32 = iota
	// Error indicates some errors occur.
	Error
)

var bufferPool = util.NewLimitedPool(512, 4096)

var (
	// ErrMetaKVMissing some keys or values are missing.
	ErrMetaKVMissing = errors.New("wrong metadata lines. some keys or values are missing")
	// ErrMessageTooLong message is too long
	ErrMessageTooLong = errors.New("message is too long")

	ErrUnsupportedCompressor = errors.New("unsupported compressor")
	ErrMetaLenError          = errors.New("meta longer than body")
	ErrMetaParseError        = errors.New("meta parse error")
)

var MaxMessageLength = 0

var magicString = []byte("PRPC")

// Compressors are compressors supported by flare. You can add customized compressor in Compressors.
var Compressors = map[CompressType]Compressor{
	CompressType_COMPRESS_TYPE_NONE: &RawDataCompressor{},
	CompressType_COMPRESS_TYPE_GZIP: &GzipCompressor{},
}

// Header magicString + bodylen + metalen
type Header [12]byte

func MagicString() []byte {
	return magicString
}
func setMagicString(h *Header) {
	copy(h[0:4], magicString)
}

func newHeader(bodyLen, metaLen uint32) *Header {
	header := Header{}
	setMagicString(&header)
	SetBodyLen(&header, bodyLen)
	SetMetaLen(&header, metaLen)
	return &header
}

func (h *Header) checkMagicString() bool {
	return bytes.Compare(h[0:4], magicString) == 0
}

func (h *Header) GetBodyLen() uint32 {
	b := h[4:8]
	return binary.BigEndian.Uint32(b)
}

func (h *Header) GetMetaLen() uint32 {
	b := h[8:]
	return binary.BigEndian.Uint32(b)
}

func SetBodyLen(h *Header, l uint32) {
	b := h[4:8]
	binary.BigEndian.PutUint32(b, l)
}

func SetMetaLen(h *Header, l uint32) {
	b := h[8:]
	binary.BigEndian.PutUint32(b, l)
}

type Message struct {
	*Header
	Meta     RpcMeta
	Metadata map[string]string
	Payload  []byte
	data     []byte
}

// NewRequestMessage creates an message for request.
func NewMessage() *Message {
	header := Header([12]byte{})
	setMagicString(&header)
	return &Message{
		Header: &header,
		Meta: RpcMeta{
			Request:  &RpcRequestMeta{},
			Response: &RpcResponseMeta{},
		},
	}
}

// NewRequestMessage creates an message for request.
func NewRequestMessage() *Message {
	header := Header([12]byte{})
	setMagicString(&header)
	m := RpcMeta{
		Request: &RpcRequestMeta{},
	}
	return &Message{
		Header: &header,
		Meta:   m,
	}
}

// NewRequestMessage creates an message for request.
func NewResponseMessage() *Message {
	header := Header([12]byte{})
	setMagicString(&header)
	m := RpcMeta{
		Response: &RpcResponseMeta{},
	}
	return &Message{
		Header: &header,
		Meta:   m,
	}
}

func (m *Message) GetServiceName() string {
	return m.Meta.Request.GetServiceName()
}

func (m *Message) SetServiceName(v string) {
	m.Meta.Request.ServiceName = v
}

func (m *Message) GetServiceMethod() string {
	return m.Meta.Request.GetMethodName()
}

func (m *Message) SetServiceMethod(v string) {
	m.Meta.Request.MethodName = v
}

func (m *Message) IsBadResponse() bool {
	return m.Meta.Response.GetErrorCode() != 0
}

func (m *Message) GetResponseErrorMsg() string {
	return m.Meta.Response.GetErrorText()
}

func (m *Message) SetResponseStatus(s int32) {
	m.Meta.Response.ErrorCode = s
}

func (m *Message) GetResponseStatus() int32 {
	return m.Meta.Response.GetErrorCode()
}

func (m *Message) SetResponseErrorMsg(s string) {
	m.Meta.Response.ErrorText = s
}

func (m *Message) GetCompressType() CompressType {
	return CompressType(m.Meta.GetCompressType())
}

func (m *Message) SetCompressType(t CompressType) {
	m.Meta.CompressType = int32(t)
}

func (m *Message) GetCorrelationId() int64 {
	return m.Meta.GetCorrelationId()
}

func (m *Message) SetCorrelationId(v int64) {
	m.Meta.CorrelationId = v
}

func (m *Message) Reset() {
	m.Payload = []byte{}
	m.data = m.data[:0]
}

// Clone clones from an message.
func (m Message) Clone() *Message {
	header := *m.Header
	c := NewMessage()
	c.Header = &header
	req := &RpcRequestMeta{}
	*req =  *m.Meta.Request
	res := &RpcResponseMeta{}
	*res = *m.Meta.Response
	c.Meta = m.Meta
	c.Meta.Request = req
	c.Meta.Response = res
	return c
}

// Read reads a message from r.
func Read(r io.Reader) (*Message, error) {
	msg := NewMessage()
	err := msg.Decode(r)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// Read reads a message from r.
func ReadResponse(r io.Reader) (*Message, error) {
	msg := NewResponseMessage()
	err := msg.Decode(r)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// Decode decodes a message from reader.
func (m *Message) Decode(r io.Reader) error {
	// validate rest length for each step?

	// parse header
	_, err := io.ReadFull(r, m.Header[:4])
	if err != nil {
		return err
	}
	if !m.Header.checkMagicString() {
		return fmt.Errorf("wrong magic string: %v", m.Header[:4])
	}

	_, err = io.ReadFull(r, m.Header[4:])
	if err != nil {
		return err
	}

	// bodyLen
	lenData := m.Header[4:8]
	bodyLen := int(binary.BigEndian.Uint32(lenData))
	metaLenData := m.Header[8:]
	metaLen := int(binary.BigEndian.Uint32(metaLenData))
	if MaxMessageLength > 0 && int(bodyLen) > MaxMessageLength {
		return ErrMessageTooLong
	}

	if metaLen > bodyLen {
		return ErrMetaLenError
	}

	totalL := int(bodyLen)
	if cap(m.data) >= totalL { // reuse data
		m.data = m.data[:totalL]
	} else {
		m.data = make([]byte, totalL)
	}
	data := m.data
	_, err = io.ReadFull(r, data)
	if err != nil {
		return err
	}

	meta := m.data[:metaLen]
	metaDecoder := codec.PBCodec{}
	err = metaDecoder.Decode(meta, &m.Meta)
	if err != nil {
		return err
	}

	attachLen := m.Meta.GetAttachmentSize()

	if attachLen > 0 {
		attach := m.data[metaLen : metaLen+int(attachLen)]
		m.Metadata, err = decodeMetadata(uint32(attachLen), attach)
	}

	m.Payload = m.data[metaLen+int(attachLen):]
	cType := m.GetCompressType()
	if cType != CompressType_COMPRESS_TYPE_NONE {
		compressor := Compressors[cType]
		if compressor == nil {
			return ErrUnsupportedCompressor
		}
		m.Payload, err = compressor.Unzip(m.Payload)
		if err != nil {
			return err
		}
	}

	return nil
}

func decodeMetadata(l uint32, data []byte) (map[string]string, error) {
	m := make(map[string]string, 10)
	n := uint32(0)
	for n < l {
		// parse one key and value
		// key
		sl := binary.BigEndian.Uint32(data[n : n+4])
		n = n + 4
		if n+sl > l-4 {
			return m, ErrMetaKVMissing
		}
		k := string(data[n : n+sl])
		n = n + sl

		// value
		sl = binary.BigEndian.Uint32(data[n : n+4])
		n = n + 4
		if n+sl > l {
			return m, ErrMetaKVMissing
		}
		v := string(data[n : n+sl])
		n = n + sl
		m[k] = v
	}

	return m, nil
}

// Encode encodes messages.
func (m Message) Encode() []byte {
	data := m.EncodeSlicePointer()
	return *data
}

// EncodeSlicePointer encodes messages as a byte slice pointer we can use pool to improve.
func (m Message) EncodeSlicePointer() *[]byte {
	var err error
	// payload
	payload := m.Payload
	if m.GetCompressType() != CompressType_COMPRESS_TYPE_NONE {
		compressor := Compressors[m.GetCompressType()]
		if compressor == nil {
			m.SetCompressType(CompressType_COMPRESS_TYPE_NONE)
		} else {
			payload, err = compressor.Zip(m.Payload)
			if err != nil {
				m.SetCompressType(CompressType_COMPRESS_TYPE_NONE)
				payload = m.Payload
			}
		}
	}

	// MetaData
	bb := bytebufferpool.Get()
	encodeMetadata(m.Metadata, bb)
	attachMent := bb.Bytes()
	m.Meta.AttachmentSize = int32(len(attachMent))

	// Meta
	encoder := codec.PBCodec{}
	mb, _ := encoder.Encode(&m.Meta)

	// Header
	bodyLen := len(mb) + len(attachMent) + len(payload)
	metaLen := len(mb)
	h := newHeader(uint32(bodyLen), uint32(metaLen))
	l := 12 + bodyLen
	metaStart := 12
	metaEnd := 12 + metaLen
	attachStart := metaEnd
	attachEnd := attachStart + len(attachMent)

	payloadStart := attachEnd

	data := bufferPool.Get(l)
	copy(*data, h[:])
	copy((*data)[metaStart:metaEnd], mb)
	if attachStart != attachEnd {
		copy((*data)[attachStart:attachEnd], attachMent)
	}
	copy((*data)[payloadStart:], payload)

	return data
}

// len,string,len,string,......
func encodeMetadata(m map[string]string, bb *bytebufferpool.ByteBuffer) {
	if len(m) == 0 {
		return
	}
	d := poolUint32Data.Get().(*[]byte)
	for k, v := range m {
		binary.BigEndian.PutUint32(*d, uint32(len(k)))
		bb.Write(*d)
		bb.Write(util.StringToSliceByte(k))
		binary.BigEndian.PutUint32(*d, uint32(len(v)))
		bb.Write(*d)
		bb.Write(util.StringToSliceByte(v))
	}
}

// PutData puts the byte slice into pool.
func PutData(data *[]byte) {
	bufferPool.Put(data)
}

// WriteTo writes message to writers.
func (m Message) WriteTo(w io.Writer) (int64, error) {
	var err error
	// payload
	payload := m.Payload
	if m.GetCompressType() != CompressType_COMPRESS_TYPE_NONE {
		compressor := Compressors[m.GetCompressType()]
		if compressor == nil {
			m.SetCompressType(CompressType_COMPRESS_TYPE_NONE)
		} else {
			payload, err = compressor.Zip(m.Payload)
			if err != nil {
				m.SetCompressType(CompressType_COMPRESS_TYPE_NONE)
				payload = m.Payload
			}
		}
	}

	// MetaData
	bb := bytebufferpool.Get()
	defer bytebufferpool.Put(bb)
	encodeMetadata(m.Metadata, bb)
	attachMent := bb.Bytes()
	m.Meta.AttachmentSize = int32(len(attachMent))

	// Meta
	encoder := codec.PBCodec{}
	mb, err := encoder.Encode(&m.Meta)
	if err != nil {
		return 0, err
	}

	// Header
	bodyLen := len(mb) + len(attachMent) + len(payload)
	metaLen := len(mb)
	h := newHeader(uint32(bodyLen), uint32(metaLen))

	nn, err := w.Write(h[:])
	n := int64(nn)
	if err != nil {
		return n, err
	}

	_, err = w.Write(mb[:])
	if err != nil {
		return n, err
	}
	if len(attachMent) > 0 {
		_, err = w.Write(attachMent[:])
		if err != nil {
			return n, err
		}
	}

	nn, err = w.Write(payload)
	return int64(nn), err
}
