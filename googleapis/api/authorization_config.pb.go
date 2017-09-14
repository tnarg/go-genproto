// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/experimental/authorization_config.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	google/api/experimental/authorization_config.proto
	google/api/experimental/experimental.proto

It has these top-level messages:
	AuthorizationConfig
	Experimental
*/
package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Configuration of authorization.
//
// This section determines the authorization provider, if unspecified, then no
// authorization check will be done.
//
// Example:
//
//     experimental:
//       authorization:
//         provider: firebaserules.googleapis.com
type AuthorizationConfig struct {
	// The name of the authorization provider, such as
	// firebaserules.googleapis.com.
	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (m *AuthorizationConfig) Reset()         { *m = AuthorizationConfig{} }
func (m *AuthorizationConfig) String() string { return proto.CompactTextString(m) }
func (*AuthorizationConfig) ProtoMessage()    {}
func (*AuthorizationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorAuthorizationConfig, []int{0}
}

func (m *AuthorizationConfig) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthorizationConfig)(nil), "google.api.AuthorizationConfig")
}
func (m *AuthorizationConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthorizationConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuthorizationConfig(dAtA, i, uint64(len(m.Provider)))
		i += copy(dAtA[i:], m.Provider)
	}
	return i, nil
}

func encodeFixed64AuthorizationConfig(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32AuthorizationConfig(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAuthorizationConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AuthorizationConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovAuthorizationConfig(uint64(l))
	}
	return n
}

func sovAuthorizationConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuthorizationConfig(x uint64) (n int) {
	return sovAuthorizationConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthorizationConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthorizationConfig
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
			return fmt.Errorf("proto: AuthorizationConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthorizationConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorizationConfig
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
				return ErrInvalidLengthAuthorizationConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthorizationConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthorizationConfig
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
func skipAuthorizationConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthorizationConfig
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
					return 0, ErrIntOverflowAuthorizationConfig
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
					return 0, ErrIntOverflowAuthorizationConfig
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
				return 0, ErrInvalidLengthAuthorizationConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuthorizationConfig
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
				next, err := skipAuthorizationConfig(dAtA[start:])
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
	ErrInvalidLengthAuthorizationConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthorizationConfig   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("google/api/experimental/authorization_config.proto", fileDescriptorAuthorizationConfig)
}

var fileDescriptorAuthorizationConfig = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xad, 0x28, 0x48, 0x2d, 0xca, 0xcc, 0x4d, 0xcd,
	0x2b, 0x49, 0xcc, 0xd1, 0x4f, 0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xca, 0xac, 0x4a, 0x2c, 0xc9, 0xcc,
	0xcf, 0x8b, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x82, 0xe8, 0xd1, 0x4b, 0x2c, 0xc8, 0x54, 0x32, 0xe4, 0x12, 0x76, 0x44, 0x56, 0xe9, 0x0c, 0x56,
	0x28, 0x24, 0xc5, 0xc5, 0x51, 0x50, 0x94, 0x5f, 0x96, 0x99, 0x92, 0x5a, 0x24, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x3b, 0xe5, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x5c, 0x7c, 0xc9, 0xf9, 0xb9, 0x7a, 0x08, 0x03, 0x9d, 0x24, 0xb0, 0x18,
	0x17, 0x00, 0xb2, 0x36, 0x80, 0x31, 0x4a, 0x17, 0xaa, 0x2e, 0x3d, 0x3f, 0x27, 0x31, 0x2f, 0x5d,
	0x2f, 0xbf, 0x28, 0x5d, 0x3f, 0x3d, 0x35, 0x0f, 0xec, 0x28, 0x7d, 0x88, 0x54, 0x62, 0x41, 0x66,
	0x31, 0xc8, 0x2f, 0xd6, 0x89, 0x05, 0x99, 0x8b, 0x98, 0x58, 0xdc, 0x1d, 0x03, 0x3c, 0x93, 0xd8,
	0xc0, 0x0a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0xba, 0xcb, 0xa0, 0xeb, 0x00, 0x00,
	0x00,
}
