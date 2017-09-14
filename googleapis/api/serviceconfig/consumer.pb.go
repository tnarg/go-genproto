// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/consumer.proto

package serviceconfig

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Supported data type of the property values
type Property_PropertyType int32

const (
	// The type is unspecified, and will result in an error.
	Property_UNSPECIFIED Property_PropertyType = 0
	// The type is `int64`.
	Property_INT64 Property_PropertyType = 1
	// The type is `bool`.
	Property_BOOL Property_PropertyType = 2
	// The type is `string`.
	Property_STRING Property_PropertyType = 3
	// The type is 'double'.
	Property_DOUBLE Property_PropertyType = 4
)

var Property_PropertyType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "INT64",
	2: "BOOL",
	3: "STRING",
	4: "DOUBLE",
}
var Property_PropertyType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"INT64":       1,
	"BOOL":        2,
	"STRING":      3,
	"DOUBLE":      4,
}

func (x Property_PropertyType) String() string {
	return proto.EnumName(Property_PropertyType_name, int32(x))
}
func (Property_PropertyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorConsumer, []int{1, 0}
}

// A descriptor for defining project properties for a service. One service may
// have many consumer projects, and the service may want to behave differently
// depending on some properties on the project. For example, a project may be
// associated with a school, or a business, or a government agency, a business
// type property on the project may affect how a service responds to the client.
// This descriptor defines which properties are allowed to be set on a project.
//
// Example:
//
//    project_properties:
//      properties:
//      - name: NO_WATERMARK
//        type: BOOL
//        description: Allows usage of the API without watermarks.
//      - name: EXTENDED_TILE_CACHE_PERIOD
//        type: INT64
type ProjectProperties struct {
	// List of per consumer project-specific properties.
	Properties []*Property `protobuf:"bytes,1,rep,name=properties" json:"properties,omitempty"`
}

func (m *ProjectProperties) Reset()                    { *m = ProjectProperties{} }
func (m *ProjectProperties) String() string            { return proto.CompactTextString(m) }
func (*ProjectProperties) ProtoMessage()               {}
func (*ProjectProperties) Descriptor() ([]byte, []int) { return fileDescriptorConsumer, []int{0} }

func (m *ProjectProperties) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

// Defines project properties.
//
// API services can define properties that can be assigned to consumer projects
// so that backends can perform response customization without having to make
// additional calls or maintain additional storage. For example, Maps API
// defines properties that controls map tile cache period, or whether to embed a
// watermark in a result.
//
// These values can be set via API producer console. Only API providers can
// define and set these properties.
type Property struct {
	// The name of the property (a.k.a key).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of this property.
	Type Property_PropertyType `protobuf:"varint,2,opt,name=type,proto3,enum=google.api.Property_PropertyType" json:"type,omitempty"`
	// The description of the property
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *Property) Reset()                    { *m = Property{} }
func (m *Property) String() string            { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()               {}
func (*Property) Descriptor() ([]byte, []int) { return fileDescriptorConsumer, []int{1} }

func (m *Property) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Property) GetType() Property_PropertyType {
	if m != nil {
		return m.Type
	}
	return Property_UNSPECIFIED
}

func (m *Property) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ProjectProperties)(nil), "google.api.ProjectProperties")
	proto.RegisterType((*Property)(nil), "google.api.Property")
	proto.RegisterEnum("google.api.Property_PropertyType", Property_PropertyType_name, Property_PropertyType_value)
}
func (m *ProjectProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectProperties) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Properties) > 0 {
		for _, msg := range m.Properties {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConsumer(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Property) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Property) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConsumer(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConsumer(dAtA, i, uint64(m.Type))
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConsumer(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	return i, nil
}

func encodeFixed64Consumer(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Consumer(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConsumer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ProjectProperties) Size() (n int) {
	var l int
	_ = l
	if len(m.Properties) > 0 {
		for _, e := range m.Properties {
			l = e.Size()
			n += 1 + l + sovConsumer(uint64(l))
		}
	}
	return n
}

func (m *Property) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConsumer(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovConsumer(uint64(m.Type))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovConsumer(uint64(l))
	}
	return n
}

func sovConsumer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConsumer(x uint64) (n int) {
	return sovConsumer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProjectProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsumer
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
			return fmt.Errorf("proto: ProjectProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumer
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
				return ErrInvalidLengthConsumer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Properties = append(m.Properties, &Property{})
			if err := m.Properties[len(m.Properties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsumer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConsumer
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
func (m *Property) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsumer
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
			return fmt.Errorf("proto: Property: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Property: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumer
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
				return ErrInvalidLengthConsumer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Property_PropertyType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumer
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
				return ErrInvalidLengthConsumer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsumer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConsumer
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
func skipConsumer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConsumer
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
					return 0, ErrIntOverflowConsumer
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
					return 0, ErrIntOverflowConsumer
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
				return 0, ErrInvalidLengthConsumer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConsumer
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
				next, err := skipConsumer(dAtA[start:])
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
	ErrInvalidLengthConsumer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConsumer   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("google/api/consumer.proto", fileDescriptorConsumer) }

var fileDescriptorConsumer = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4f, 0xfa, 0x40,
	0x10, 0xc5, 0xff, 0x0b, 0xfd, 0x13, 0x18, 0x14, 0xeb, 0xc6, 0x43, 0xbd, 0x90, 0xca, 0x89, 0x53,
	0x9b, 0x20, 0x7a, 0xf1, 0x56, 0xa8, 0xa6, 0x09, 0x81, 0xa6, 0xc0, 0xc5, 0x5b, 0xad, 0xe3, 0x66,
	0x0d, 0x74, 0x36, 0xdb, 0x6a, 0xc2, 0x07, 0x34, 0xf1, 0xe8, 0x47, 0x30, 0x7c, 0x12, 0x43, 0xa9,
	0x58, 0x13, 0x6f, 0x6f, 0xf6, 0xf7, 0xde, 0x64, 0xf3, 0x06, 0xce, 0x05, 0x91, 0x58, 0xa1, 0x1b,
	0x2b, 0xe9, 0x26, 0x94, 0x66, 0x2f, 0x6b, 0xd4, 0x8e, 0xd2, 0x94, 0x13, 0x87, 0x3d, 0x72, 0x62,
	0x25, 0x7b, 0x01, 0x9c, 0x86, 0x9a, 0x9e, 0x31, 0xc9, 0x43, 0x4d, 0x0a, 0x75, 0x2e, 0x31, 0xe3,
	0x43, 0x00, 0x75, 0x98, 0x2c, 0x66, 0xd7, 0xfb, 0xed, 0xc1, 0x99, 0xf3, 0x93, 0x72, 0x4a, 0xef,
	0x26, 0xaa, 0xf8, 0x7a, 0x6f, 0x0c, 0x9a, 0xdf, 0x80, 0x73, 0x30, 0xd2, 0x78, 0x8d, 0x16, 0xb3,
	0x59, 0xbf, 0x15, 0x15, 0x9a, 0x5f, 0x81, 0x91, 0x6f, 0x14, 0x5a, 0x35, 0x9b, 0xf5, 0x3b, 0x83,
	0x8b, 0xbf, 0x16, 0x1e, 0xc4, 0x62, 0xa3, 0x30, 0x2a, 0xec, 0xdc, 0x86, 0xf6, 0x23, 0x66, 0x89,
	0x96, 0x2a, 0x97, 0x94, 0x5a, 0xf5, 0x62, 0x63, 0xf5, 0xa9, 0x37, 0x81, 0xa3, 0x6a, 0x8e, 0x9f,
	0x40, 0x7b, 0x39, 0x9d, 0x87, 0xfe, 0x28, 0xb8, 0x0d, 0xfc, 0xb1, 0xf9, 0x8f, 0xb7, 0xe0, 0x7f,
	0x30, 0x5d, 0x5c, 0x0f, 0x4d, 0xc6, 0x9b, 0x60, 0x78, 0xb3, 0xd9, 0xc4, 0xac, 0x71, 0x80, 0xc6,
	0x7c, 0x11, 0x05, 0xd3, 0x3b, 0xb3, 0xbe, 0xd3, 0xe3, 0xd9, 0xd2, 0x9b, 0xf8, 0xa6, 0xe1, 0xe5,
	0xef, 0xdb, 0x2e, 0xfb, 0xd8, 0x76, 0xd9, 0xe7, 0xb6, 0xcb, 0xa0, 0x93, 0xd0, 0xba, 0xf2, 0x53,
	0xef, 0x78, 0x54, 0x96, 0x19, 0xee, 0xba, 0x0c, 0xd9, 0xbd, 0x5f, 0x42, 0x41, 0xab, 0x38, 0x15,
	0x0e, 0x69, 0xe1, 0x0a, 0x4c, 0x8b, 0xa6, 0xdd, 0x3d, 0x8a, 0x95, 0xcc, 0x8a, 0x3b, 0x64, 0xa8,
	0x5f, 0x65, 0x82, 0x09, 0xa5, 0x4f, 0x52, 0xdc, 0xfc, 0x9a, 0x1e, 0x1a, 0x45, 0xe2, 0xf2, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x03, 0x97, 0xf0, 0xb8, 0x01, 0x00, 0x00,
}
