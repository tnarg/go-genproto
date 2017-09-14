// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/servicecontrol/v1/log_entry.proto

package servicecontrol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_logging_type "google.golang.org/genproto/googleapis/logging/type"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"
import google_protobuf3 "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An individual log entry.
type LogEntry struct {
	// Required. The log to which this log entry belongs. Examples: `"syslog"`,
	// `"book_log"`.
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// The time the event described by the log entry occurred. If
	// omitted, defaults to operation start time.
	Timestamp *google_protobuf3.Timestamp `protobuf:"bytes,11,opt,name=timestamp" json:"timestamp,omitempty"`
	// The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity google_logging_type.LogSeverity `protobuf:"varint,12,opt,name=severity,proto3,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// A unique ID for the log entry used for deduplication. If omitted,
	// the implementation will generate one based on operation_id.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId,proto3" json:"insert_id,omitempty"`
	// A set of user-defined (key, value) data that provides additional
	// information about the log entry.
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The log entry payload, which can be one of multiple types.
	//
	// Types that are valid to be assigned to Payload:
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_StructPayload
	Payload isLogEntry_Payload `protobuf_oneof:"payload"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptorLogEntry, []int{0} }

type isLogEntry_Payload interface {
	isLogEntry_Payload()
	MarshalTo([]byte) (int, error)
	Size() int
}

type LogEntry_ProtoPayload struct {
	ProtoPayload *google_protobuf1.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,oneof"`
}
type LogEntry_TextPayload struct {
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,proto3,oneof"`
}
type LogEntry_StructPayload struct {
	StructPayload *google_protobuf2.Struct `protobuf:"bytes,6,opt,name=struct_payload,json=structPayload,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload()  {}
func (*LogEntry_TextPayload) isLogEntry_Payload()   {}
func (*LogEntry_StructPayload) isLogEntry_Payload() {}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *LogEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogEntry) GetTimestamp() *google_protobuf3.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetSeverity() google_logging_type.LogSeverity {
	if m != nil {
		return m.Severity
	}
	return google_logging_type.LogSeverity_DEFAULT
}

func (m *LogEntry) GetInsertId() string {
	if m != nil {
		return m.InsertId
	}
	return ""
}

func (m *LogEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LogEntry) GetProtoPayload() *google_protobuf1.Any {
	if x, ok := m.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (m *LogEntry) GetTextPayload() string {
	if x, ok := m.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (m *LogEntry) GetStructPayload() *google_protobuf2.Struct {
	if x, ok := m.GetPayload().(*LogEntry_StructPayload); ok {
		return x.StructPayload
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogEntry_OneofMarshaler, _LogEntry_OneofUnmarshaler, _LogEntry_OneofSizer, []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_StructPayload)(nil),
	}
}

func _LogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProtoPayload); err != nil {
			return err
		}
	case *LogEntry_TextPayload:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.TextPayload)
	case *LogEntry_StructPayload:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StructPayload); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LogEntry.Payload has unexpected type %T", x)
	}
	return nil
}

func _LogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEntry)
	switch tag {
	case 2: // payload.proto_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Any)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_ProtoPayload{msg}
		return true, err
	case 3: // payload.text_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &LogEntry_TextPayload{x}
		return true, err
	case 6: // payload.struct_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf2.Struct)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_StructPayload{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		s := proto.Size(x.ProtoPayload)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LogEntry_TextPayload:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TextPayload)))
		n += len(x.TextPayload)
	case *LogEntry_StructPayload:
		s := proto.Size(x.StructPayload)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "google.api.servicecontrol.v1.LogEntry")
}
func (m *LogEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		nn1, err := m.Payload.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if len(m.InsertId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLogEntry(dAtA, i, uint64(len(m.InsertId)))
		i += copy(dAtA[i:], m.InsertId)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintLogEntry(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Timestamp != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintLogEntry(dAtA, i, uint64(m.Timestamp.Size()))
		n2, err := m.Timestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Severity != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintLogEntry(dAtA, i, uint64(m.Severity))
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x6a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovLogEntry(uint64(len(k))) + 1 + len(v) + sovLogEntry(uint64(len(v)))
			i = encodeVarintLogEntry(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintLogEntry(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintLogEntry(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *LogEntry_ProtoPayload) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ProtoPayload != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogEntry(dAtA, i, uint64(m.ProtoPayload.Size()))
		n3, err := m.ProtoPayload.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *LogEntry_TextPayload) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x1a
	i++
	i = encodeVarintLogEntry(dAtA, i, uint64(len(m.TextPayload)))
	i += copy(dAtA[i:], m.TextPayload)
	return i, nil
}
func (m *LogEntry_StructPayload) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.StructPayload != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintLogEntry(dAtA, i, uint64(m.StructPayload.Size()))
		n4, err := m.StructPayload.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func encodeFixed64LogEntry(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32LogEntry(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLogEntry(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LogEntry) Size() (n int) {
	var l int
	_ = l
	if m.Payload != nil {
		n += m.Payload.Size()
	}
	l = len(m.InsertId)
	if l > 0 {
		n += 1 + l + sovLogEntry(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLogEntry(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovLogEntry(uint64(l))
	}
	if m.Severity != 0 {
		n += 1 + sovLogEntry(uint64(m.Severity))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovLogEntry(uint64(len(k))) + 1 + len(v) + sovLogEntry(uint64(len(v)))
			n += mapEntrySize + 1 + sovLogEntry(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *LogEntry_ProtoPayload) Size() (n int) {
	var l int
	_ = l
	if m.ProtoPayload != nil {
		l = m.ProtoPayload.Size()
		n += 1 + l + sovLogEntry(uint64(l))
	}
	return n
}
func (m *LogEntry_TextPayload) Size() (n int) {
	var l int
	_ = l
	l = len(m.TextPayload)
	n += 1 + l + sovLogEntry(uint64(l))
	return n
}
func (m *LogEntry_StructPayload) Size() (n int) {
	var l int
	_ = l
	if m.StructPayload != nil {
		l = m.StructPayload.Size()
		n += 1 + l + sovLogEntry(uint64(l))
	}
	return n
}

func sovLogEntry(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLogEntry(x uint64) (n int) {
	return sovLogEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LogEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoPayload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &google_protobuf1.Any{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Payload = &LogEntry_ProtoPayload{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TextPayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = &LogEntry_TextPayload{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsertId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InsertId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StructPayload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &google_protobuf2.Struct{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Payload = &LogEntry_StructPayload{v}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &google_protobuf3.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			m.Severity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Severity |= (google_logging_type.LogSeverity(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLogEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLogEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthLogEntry
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLogEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthLogEntry
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipLogEntry(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthLogEntry
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogEntry
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
func skipLogEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogEntry
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
					return 0, ErrIntOverflowLogEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogEntry
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLogEntry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLogEntry
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLogEntry(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLogEntry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogEntry   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("google/api/servicecontrol/v1/log_entry.proto", fileDescriptorLogEntry)
}

var fileDescriptorLogEntry = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0xeb, 0xed, 0x52, 0x1a, 0xa7, 0x5d, 0x21, 0x6b, 0x25, 0x42, 0xa8, 0x4a, 0x04, 0x12,
	0xea, 0x01, 0x39, 0xda, 0x72, 0x59, 0x58, 0x0e, 0x50, 0x09, 0xa9, 0xac, 0x7a, 0x58, 0x65, 0x39,
	0x71, 0xa9, 0xdc, 0xd6, 0x58, 0x16, 0xae, 0x1d, 0x25, 0x6e, 0x44, 0x9e, 0x81, 0x17, 0xe3, 0xc8,
	0x23, 0xa0, 0xde, 0x78, 0x0b, 0x14, 0x7f, 0x84, 0xfd, 0x40, 0xbd, 0x79, 0x66, 0x7e, 0xff, 0xbf,
	0x67, 0x3c, 0x09, 0x7c, 0xc5, 0x94, 0x62, 0x82, 0xa6, 0x24, 0xe7, 0x69, 0x49, 0x8b, 0x8a, 0xaf,
	0xe9, 0x5a, 0x49, 0x5d, 0x28, 0x91, 0x56, 0x67, 0xa9, 0x50, 0x6c, 0x49, 0xa5, 0x2e, 0x6a, 0x9c,
	0x17, 0x4a, 0x2b, 0x34, 0xb2, 0x34, 0x26, 0x39, 0xc7, 0xb7, 0x69, 0x5c, 0x9d, 0xc5, 0xa3, 0x1b,
	0x5e, 0x44, 0x4a, 0xa5, 0x89, 0xe6, 0x4a, 0x96, 0x56, 0x1b, 0xbf, 0x74, 0x55, 0xa1, 0x18, 0xe3,
	0x92, 0xa5, 0xba, 0xce, 0x4d, 0xb0, 0x2c, 0x69, 0x45, 0x0b, 0xae, 0xdd, 0x1d, 0xf1, 0x13, 0xc7,
	0x99, 0x68, 0xb5, 0xfb, 0x9a, 0x12, 0xe9, 0x4b, 0xa3, 0xbb, 0xa5, 0x52, 0x17, 0xbb, 0xb5, 0x76,
	0xd5, 0x67, 0x77, 0xab, 0x9a, 0x6f, 0x69, 0xa9, 0xc9, 0x36, 0xb7, 0xc0, 0xf3, 0x3f, 0x5d, 0xd8,
	0x5f, 0x28, 0xf6, 0xb1, 0x19, 0x08, 0x5d, 0xc0, 0xa1, 0xc9, 0x2e, 0x73, 0x52, 0x0b, 0x45, 0x36,
	0xd1, 0x51, 0x02, 0x26, 0xe1, 0xf4, 0x14, 0xbb, 0x11, 0xbd, 0x0b, 0xfe, 0x20, 0xeb, 0x79, 0x27,
	0x1b, 0x98, 0xf8, 0xca, 0xb2, 0xe8, 0x05, 0x1c, 0x68, 0xfa, 0x5d, 0xb7, 0xda, 0x6e, 0x02, 0x26,
	0xc1, 0xbc, 0x93, 0x85, 0x4d, 0xd6, 0x43, 0x4f, 0x61, 0xc0, 0x65, 0x49, 0x0b, 0xbd, 0xe4, 0x9b,
	0xe8, 0xb8, 0x21, 0xb2, 0xbe, 0x4d, 0x7c, 0xda, 0xa0, 0xf7, 0xf0, 0xc4, 0x36, 0xdf, 0x7a, 0xf4,
	0xcc, 0xfd, 0x8f, 0xef, 0xdd, 0x7f, 0x6d, 0xb0, 0x79, 0x27, 0x1b, 0x5a, 0x81, 0xb7, 0x47, 0xf0,
	0x58, 0x92, 0x2d, 0x8d, 0xa0, 0x71, 0x36, 0x67, 0x74, 0x0e, 0x83, 0x76, 0xe8, 0x28, 0x34, 0x86,
	0xf1, 0x3d, 0xc3, 0xcf, 0x9e, 0xc8, 0xfe, 0xc1, 0xe8, 0x1d, 0xec, 0xfb, 0x3d, 0x44, 0x83, 0x04,
	0x4c, 0x4e, 0xa6, 0x89, 0x17, 0xba, 0x85, 0xe1, 0x66, 0x61, 0x78, 0xa1, 0xd8, 0xb5, 0xe3, 0xb2,
	0x56, 0x81, 0x2e, 0x61, 0x4f, 0x90, 0x15, 0x15, 0x65, 0x34, 0x4c, 0xba, 0x93, 0x70, 0x3a, 0xc5,
	0x87, 0x3e, 0x14, 0xec, 0x97, 0x80, 0x17, 0x46, 0x64, 0xce, 0x99, 0x73, 0x88, 0xdf, 0xc0, 0xf0,
	0x46, 0x1a, 0x3d, 0x82, 0xdd, 0x6f, 0xb4, 0x8e, 0x80, 0x99, 0xb2, 0x39, 0xa2, 0x53, 0xf8, 0xa0,
	0x22, 0x62, 0x47, 0xcd, 0xc6, 0x82, 0xcc, 0x06, 0x6f, 0x8f, 0xce, 0xc1, 0x2c, 0x80, 0x0f, 0xdd,
	0x6b, 0xce, 0x7e, 0x80, 0x9f, 0xfb, 0x31, 0xf8, 0xb5, 0x1f, 0x83, 0xdf, 0xfb, 0x31, 0x80, 0xc9,
	0x5a, 0x6d, 0x0f, 0xf6, 0x34, 0x1b, 0xfa, 0xa6, 0xae, 0xcc, 0xa2, 0xc1, 0x97, 0x4b, 0x87, 0x33,
	0x25, 0x88, 0x64, 0x58, 0x15, 0x2c, 0x65, 0x54, 0x9a, 0x57, 0x4c, 0x6d, 0x89, 0xe4, 0xbc, 0xfc,
	0xff, 0x8f, 0x73, 0x71, 0x3b, 0xb3, 0xea, 0x19, 0xd9, 0xeb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x5b, 0xda, 0xbc, 0xc8, 0x6e, 0x03, 0x00, 0x00,
}
