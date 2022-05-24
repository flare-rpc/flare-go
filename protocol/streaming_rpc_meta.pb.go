// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: streaming_rpc_meta.proto

package protocol

import (
	fmt "fmt"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FrameType int32

const (
	FrameType_FRAME_TYPE_UNKNOWN  FrameType = 0
	FrameType_FRAME_TYPE_RST      FrameType = 1
	FrameType_FRAME_TYPE_CLOSE    FrameType = 2
	FrameType_FRAME_TYPE_DATA     FrameType = 3
	FrameType_FRAME_TYPE_FEEDBACK FrameType = 4
)

var FrameType_name = map[int32]string{
	0: "FRAME_TYPE_UNKNOWN",
	1: "FRAME_TYPE_RST",
	2: "FRAME_TYPE_CLOSE",
	3: "FRAME_TYPE_DATA",
	4: "FRAME_TYPE_FEEDBACK",
}

var FrameType_value = map[string]int32{
	"FRAME_TYPE_UNKNOWN":  0,
	"FRAME_TYPE_RST":      1,
	"FRAME_TYPE_CLOSE":    2,
	"FRAME_TYPE_DATA":     3,
	"FRAME_TYPE_FEEDBACK": 4,
}

func (x FrameType) Enum() *FrameType {
	p := new(FrameType)
	*p = x
	return p
}

func (x FrameType) String() string {
	return proto.EnumName(FrameType_name, int32(x))
}

func (x *FrameType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FrameType_value, data, "FrameType")
	if err != nil {
		return err
	}
	*x = FrameType(value)
	return nil
}

func (FrameType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4b3a27f2b175701, []int{0}
}

type StreamSettings struct {
	StreamId     int64 `protobuf:"varint,1,req,name=stream_id,json=streamId" json:"stream_id"`
	NeedFeedback *bool `protobuf:"varint,2,opt,name=need_feedback,json=needFeedback,def=0" json:"need_feedback,omitempty"`
	Writable     *bool `protobuf:"varint,3,opt,name=writable,def=0" json:"writable,omitempty"`
}

func (m *StreamSettings) Reset()         { *m = StreamSettings{} }
func (m *StreamSettings) String() string { return proto.CompactTextString(m) }
func (*StreamSettings) ProtoMessage()    {}
func (*StreamSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b3a27f2b175701, []int{0}
}
func (m *StreamSettings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamSettings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamSettings.Merge(m, src)
}
func (m *StreamSettings) XXX_Size() int {
	return m.Size()
}
func (m *StreamSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamSettings.DiscardUnknown(m)
}

var xxx_messageInfo_StreamSettings proto.InternalMessageInfo

const Default_StreamSettings_NeedFeedback bool = false
const Default_StreamSettings_Writable bool = false

func (m *StreamSettings) GetStreamId() int64 {
	if m != nil {
		return m.StreamId
	}
	return 0
}

func (m *StreamSettings) GetNeedFeedback() bool {
	if m != nil && m.NeedFeedback != nil {
		return *m.NeedFeedback
	}
	return Default_StreamSettings_NeedFeedback
}

func (m *StreamSettings) GetWritable() bool {
	if m != nil && m.Writable != nil {
		return *m.Writable
	}
	return Default_StreamSettings_Writable
}

type StreamFrameMeta struct {
	StreamId        int64     `protobuf:"varint,1,req,name=stream_id,json=streamId" json:"stream_id"`
	SourceStreamId  int64     `protobuf:"varint,2,opt,name=source_stream_id,json=sourceStreamId" json:"source_stream_id"`
	FrameType       FrameType `protobuf:"varint,3,opt,name=frame_type,json=frameType,enum=flare.rpc.FrameType" json:"frame_type"`
	HasContinuation bool      `protobuf:"varint,4,opt,name=has_continuation,json=hasContinuation" json:"has_continuation"`
	Feedback        *Feedback `protobuf:"bytes,5,opt,name=feedback" json:"feedback,omitempty"`
}

func (m *StreamFrameMeta) Reset()         { *m = StreamFrameMeta{} }
func (m *StreamFrameMeta) String() string { return proto.CompactTextString(m) }
func (*StreamFrameMeta) ProtoMessage()    {}
func (*StreamFrameMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b3a27f2b175701, []int{1}
}
func (m *StreamFrameMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamFrameMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamFrameMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamFrameMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamFrameMeta.Merge(m, src)
}
func (m *StreamFrameMeta) XXX_Size() int {
	return m.Size()
}
func (m *StreamFrameMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamFrameMeta.DiscardUnknown(m)
}

var xxx_messageInfo_StreamFrameMeta proto.InternalMessageInfo

func (m *StreamFrameMeta) GetStreamId() int64 {
	if m != nil {
		return m.StreamId
	}
	return 0
}

func (m *StreamFrameMeta) GetSourceStreamId() int64 {
	if m != nil {
		return m.SourceStreamId
	}
	return 0
}

func (m *StreamFrameMeta) GetFrameType() FrameType {
	if m != nil {
		return m.FrameType
	}
	return FrameType_FRAME_TYPE_UNKNOWN
}

func (m *StreamFrameMeta) GetHasContinuation() bool {
	if m != nil {
		return m.HasContinuation
	}
	return false
}

func (m *StreamFrameMeta) GetFeedback() *Feedback {
	if m != nil {
		return m.Feedback
	}
	return nil
}

type Feedback struct {
	ConsumedSize int64 `protobuf:"varint,1,opt,name=consumed_size,json=consumedSize" json:"consumed_size"`
}

func (m *Feedback) Reset()         { *m = Feedback{} }
func (m *Feedback) String() string { return proto.CompactTextString(m) }
func (*Feedback) ProtoMessage()    {}
func (*Feedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b3a27f2b175701, []int{2}
}
func (m *Feedback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Feedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Feedback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Feedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feedback.Merge(m, src)
}
func (m *Feedback) XXX_Size() int {
	return m.Size()
}
func (m *Feedback) XXX_DiscardUnknown() {
	xxx_messageInfo_Feedback.DiscardUnknown(m)
}

var xxx_messageInfo_Feedback proto.InternalMessageInfo

func (m *Feedback) GetConsumedSize() int64 {
	if m != nil {
		return m.ConsumedSize
	}
	return 0
}

func init() {
	proto.RegisterEnum("flare.rpc.FrameType", FrameType_name, FrameType_value)
	proto.RegisterType((*StreamSettings)(nil), "flare.rpc.StreamSettings")
	proto.RegisterType((*StreamFrameMeta)(nil), "flare.rpc.StreamFrameMeta")
	proto.RegisterType((*Feedback)(nil), "flare.rpc.Feedback")
}

func init() { proto.RegisterFile("streaming_rpc_meta.proto", fileDescriptor_b4b3a27f2b175701) }

var fileDescriptor_b4b3a27f2b175701 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xbd, 0x49, 0x2a, 0x39, 0x43, 0x9b, 0x98, 0x4d, 0x05, 0x3e, 0x99, 0x34, 0x12, 0x52,
	0xa8, 0x84, 0x23, 0x55, 0xe2, 0x00, 0xb7, 0x24, 0x75, 0x04, 0x2a, 0x4d, 0x2b, 0x3b, 0x08, 0xc1,
	0x65, 0xb5, 0x59, 0xaf, 0x13, 0x8b, 0xd8, 0x6b, 0xad, 0x37, 0x42, 0x2d, 0x27, 0x78, 0x02, 0x1e,
	0xab, 0xc7, 0x1e, 0x39, 0x21, 0x94, 0xbc, 0x08, 0x72, 0x1c, 0xbb, 0x3e, 0xf6, 0xb6, 0xfa, 0xff,
	0x6f, 0x76, 0xfe, 0x19, 0x0d, 0x98, 0xa9, 0x92, 0x9c, 0x46, 0x61, 0xbc, 0x20, 0x32, 0x61, 0x24,
	0xe2, 0x8a, 0xda, 0x89, 0x14, 0x4a, 0xe0, 0x66, 0xb0, 0xa2, 0x92, 0xdb, 0x32, 0x61, 0xbd, 0x5f,
	0x08, 0x5a, 0xde, 0x8e, 0xf3, 0xb8, 0x52, 0x61, 0xbc, 0x48, 0xf1, 0x09, 0x34, 0xf3, 0x4a, 0x12,
	0xfa, 0x26, 0xea, 0xd6, 0xfa, 0xf5, 0x51, 0xe3, 0xee, 0xef, 0x0b, 0xcd, 0xd5, 0x73, 0xf9, 0x83,
	0x8f, 0x4f, 0xe1, 0x28, 0xe6, 0xdc, 0x27, 0x01, 0xe7, 0xfe, 0x9c, 0xb2, 0x6f, 0x66, 0xad, 0x8b,
	0xfa, 0xfa, 0xbb, 0x83, 0x80, 0xae, 0x52, 0xee, 0x1e, 0x66, 0xde, 0x64, 0x6f, 0xe1, 0x13, 0xd0,
	0xbf, 0xcb, 0x50, 0xd1, 0xf9, 0x8a, 0x9b, 0xf5, 0x2a, 0x56, 0xca, 0xbd, 0x9f, 0x35, 0x68, 0xe7,
	0x21, 0x26, 0x92, 0x46, 0xfc, 0x92, 0x2b, 0xfa, 0x98, 0x14, 0x36, 0x18, 0xa9, 0x58, 0x4b, 0xc6,
	0xc9, 0x03, 0x99, 0x05, 0x29, 0xc8, 0x56, 0xee, 0x7a, 0x05, 0xff, 0x16, 0x20, 0xc8, 0xfe, 0x27,
	0xea, 0x26, 0xc9, 0xb3, 0xb4, 0xce, 0x8e, 0xed, 0x72, 0x17, 0xf6, 0xae, 0xf9, 0xec, 0x26, 0xe1,
	0xfb, 0xfa, 0x66, 0x50, 0x08, 0x78, 0x00, 0xc6, 0x92, 0xa6, 0x84, 0x89, 0x58, 0x85, 0xf1, 0x9a,
	0xaa, 0x50, 0xc4, 0x66, 0x23, 0x1b, 0x66, 0x8f, 0xb6, 0x97, 0x34, 0x1d, 0x57, 0x4c, 0x3c, 0x00,
	0xbd, 0x5c, 0xce, 0x41, 0x17, 0xf5, 0x9f, 0x9c, 0x75, 0xaa, 0x9d, 0xf6, 0x96, 0x5b, 0x42, 0xbd,
	0x37, 0xa0, 0x97, 0x2b, 0x7b, 0x05, 0x47, 0x4c, 0xc4, 0xe9, 0x3a, 0xe2, 0x3e, 0x49, 0xc3, 0x5b,
	0x6e, 0xa2, 0xca, 0x54, 0x87, 0x85, 0xe5, 0x85, 0xb7, 0xfc, 0xf4, 0x07, 0x34, 0xcb, 0xd8, 0xf8,
	0x19, 0xe0, 0x89, 0x3b, 0xbc, 0x74, 0xc8, 0xec, 0xcb, 0xb5, 0x43, 0x3e, 0x4d, 0x2f, 0xa6, 0x57,
	0x9f, 0xa7, 0x86, 0x86, 0x31, 0xb4, 0x2a, 0xba, 0xeb, 0xcd, 0x0c, 0x84, 0x8f, 0xc1, 0xa8, 0x68,
	0xe3, 0x8f, 0x57, 0x9e, 0x63, 0xd4, 0x70, 0x07, 0xda, 0x15, 0xf5, 0x7c, 0x38, 0x1b, 0x1a, 0x75,
	0xfc, 0x1c, 0x3a, 0x15, 0x71, 0xe2, 0x38, 0xe7, 0xa3, 0xe1, 0xf8, 0xc2, 0x68, 0x8c, 0xd8, 0xdd,
	0xc6, 0x42, 0xf7, 0x1b, 0x0b, 0xfd, 0xdb, 0x58, 0xe8, 0xf7, 0xd6, 0xd2, 0xee, 0xb7, 0x96, 0xf6,
	0x67, 0x6b, 0x69, 0x59, 0xfe, 0xe8, 0x61, 0xde, 0xd1, 0x53, 0xaf, 0x38, 0x45, 0x37, 0x61, 0xd7,
	0xd9, 0x0d, 0xbe, 0x47, 0x5f, 0x5f, 0x2e, 0x42, 0xb5, 0x5c, 0xcf, 0x6d, 0x26, 0xa2, 0xc1, 0x0e,
	0x7d, 0x2d, 0x13, 0x96, 0xbf, 0x16, 0x62, 0xb0, 0xbb, 0x54, 0x26, 0x56, 0xff, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x05, 0x55, 0x65, 0x73, 0xc7, 0x02, 0x00, 0x00,
}

func (m *StreamSettings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamSettings) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamSettings) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Writable != nil {
		i--
		if *m.Writable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.NeedFeedback != nil {
		i--
		if *m.NeedFeedback {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	i = encodeVarintStreamingRpcMeta(dAtA, i, uint64(m.StreamId))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *StreamFrameMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamFrameMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamFrameMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Feedback != nil {
		{
			size, err := m.Feedback.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStreamingRpcMeta(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	i--
	if m.HasContinuation {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x20
	i = encodeVarintStreamingRpcMeta(dAtA, i, uint64(m.FrameType))
	i--
	dAtA[i] = 0x18
	i = encodeVarintStreamingRpcMeta(dAtA, i, uint64(m.SourceStreamId))
	i--
	dAtA[i] = 0x10
	i = encodeVarintStreamingRpcMeta(dAtA, i, uint64(m.StreamId))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *Feedback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Feedback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Feedback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintStreamingRpcMeta(dAtA, i, uint64(m.ConsumedSize))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintStreamingRpcMeta(dAtA []byte, offset int, v uint64) int {
	offset -= sovStreamingRpcMeta(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StreamSettings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovStreamingRpcMeta(uint64(m.StreamId))
	if m.NeedFeedback != nil {
		n += 2
	}
	if m.Writable != nil {
		n += 2
	}
	return n
}

func (m *StreamFrameMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovStreamingRpcMeta(uint64(m.StreamId))
	n += 1 + sovStreamingRpcMeta(uint64(m.SourceStreamId))
	n += 1 + sovStreamingRpcMeta(uint64(m.FrameType))
	n += 2
	if m.Feedback != nil {
		l = m.Feedback.Size()
		n += 1 + l + sovStreamingRpcMeta(uint64(l))
	}
	return n
}

func (m *Feedback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovStreamingRpcMeta(uint64(m.ConsumedSize))
	return n
}

func sovStreamingRpcMeta(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStreamingRpcMeta(x uint64) (n int) {
	return sovStreamingRpcMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamSettings) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreamingRpcMeta
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamSettings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamSettings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamId", wireType)
			}
			m.StreamId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NeedFeedback", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.NeedFeedback = &b
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Writable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Writable = &b
		default:
			iNdEx = preIndex
			skippy, err := skipStreamingRpcMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreamingRpcMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("stream_id")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StreamFrameMeta) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreamingRpcMeta
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StreamFrameMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamFrameMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamId", wireType)
			}
			m.StreamId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceStreamId", wireType)
			}
			m.SourceStreamId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SourceStreamId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrameType", wireType)
			}
			m.FrameType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FrameType |= FrameType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasContinuation", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasContinuation = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feedback", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStreamingRpcMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreamingRpcMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Feedback == nil {
				m.Feedback = &Feedback{}
			}
			if err := m.Feedback.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStreamingRpcMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreamingRpcMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("stream_id")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Feedback) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreamingRpcMeta
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Feedback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Feedback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsumedSize", wireType)
			}
			m.ConsumedSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConsumedSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStreamingRpcMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreamingRpcMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStreamingRpcMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStreamingRpcMeta
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStreamingRpcMeta
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStreamingRpcMeta
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStreamingRpcMeta
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStreamingRpcMeta
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStreamingRpcMeta        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStreamingRpcMeta          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStreamingRpcMeta = fmt.Errorf("proto: unexpected end of group")
)