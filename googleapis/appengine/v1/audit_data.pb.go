// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/appengine/v1/audit_data.proto

package appengine

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/iam/v1"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// App Engine admin service audit log.
type AuditData struct {
	// Detailed information about methods that require it. Does not include
	// simple Get, List or Delete methods because all significant information
	// (resource name, number of returned elements for List operations) is already
	// included in parent audit log message.
	//
	// Types that are valid to be assigned to Method:
	//	*AuditData_UpdateService
	//	*AuditData_CreateVersion
	Method isAuditData_Method `protobuf_oneof:"method"`
}

func (m *AuditData) Reset()                    { *m = AuditData{} }
func (m *AuditData) String() string            { return proto.CompactTextString(m) }
func (*AuditData) ProtoMessage()               {}
func (*AuditData) Descriptor() ([]byte, []int) { return fileDescriptorAuditData, []int{0} }

type isAuditData_Method interface {
	isAuditData_Method()
	MarshalTo([]byte) (int, error)
	Size() int
}

type AuditData_UpdateService struct {
	UpdateService *UpdateServiceMethod `protobuf:"bytes,1,opt,name=update_service,json=updateService,oneof"`
}
type AuditData_CreateVersion struct {
	CreateVersion *CreateVersionMethod `protobuf:"bytes,2,opt,name=create_version,json=createVersion,oneof"`
}

func (*AuditData_UpdateService) isAuditData_Method() {}
func (*AuditData_CreateVersion) isAuditData_Method() {}

func (m *AuditData) GetMethod() isAuditData_Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *AuditData) GetUpdateService() *UpdateServiceMethod {
	if x, ok := m.GetMethod().(*AuditData_UpdateService); ok {
		return x.UpdateService
	}
	return nil
}

func (m *AuditData) GetCreateVersion() *CreateVersionMethod {
	if x, ok := m.GetMethod().(*AuditData_CreateVersion); ok {
		return x.CreateVersion
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AuditData) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AuditData_OneofMarshaler, _AuditData_OneofUnmarshaler, _AuditData_OneofSizer, []interface{}{
		(*AuditData_UpdateService)(nil),
		(*AuditData_CreateVersion)(nil),
	}
}

func _AuditData_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AuditData)
	// method
	switch x := m.Method.(type) {
	case *AuditData_UpdateService:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateService); err != nil {
			return err
		}
	case *AuditData_CreateVersion:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateVersion); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AuditData.Method has unexpected type %T", x)
	}
	return nil
}

func _AuditData_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AuditData)
	switch tag {
	case 1: // method.update_service
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateServiceMethod)
		err := b.DecodeMessage(msg)
		m.Method = &AuditData_UpdateService{msg}
		return true, err
	case 2: // method.create_version
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateVersionMethod)
		err := b.DecodeMessage(msg)
		m.Method = &AuditData_CreateVersion{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AuditData_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AuditData)
	// method
	switch x := m.Method.(type) {
	case *AuditData_UpdateService:
		s := proto.Size(x.UpdateService)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuditData_CreateVersion:
		s := proto.Size(x.CreateVersion)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Detailed information about UpdateService call.
type UpdateServiceMethod struct {
	// Update service request.
	Request *UpdateServiceRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
}

func (m *UpdateServiceMethod) Reset()                    { *m = UpdateServiceMethod{} }
func (m *UpdateServiceMethod) String() string            { return proto.CompactTextString(m) }
func (*UpdateServiceMethod) ProtoMessage()               {}
func (*UpdateServiceMethod) Descriptor() ([]byte, []int) { return fileDescriptorAuditData, []int{1} }

func (m *UpdateServiceMethod) GetRequest() *UpdateServiceRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

// Detailed information about CreateVersion call.
type CreateVersionMethod struct {
	// Create version request.
	Request *CreateVersionRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
}

func (m *CreateVersionMethod) Reset()                    { *m = CreateVersionMethod{} }
func (m *CreateVersionMethod) String() string            { return proto.CompactTextString(m) }
func (*CreateVersionMethod) ProtoMessage()               {}
func (*CreateVersionMethod) Descriptor() ([]byte, []int) { return fileDescriptorAuditData, []int{2} }

func (m *CreateVersionMethod) GetRequest() *CreateVersionRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*AuditData)(nil), "google.appengine.v1.AuditData")
	proto.RegisterType((*UpdateServiceMethod)(nil), "google.appengine.v1.UpdateServiceMethod")
	proto.RegisterType((*CreateVersionMethod)(nil), "google.appengine.v1.CreateVersionMethod")
}
func (m *AuditData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuditData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Method != nil {
		nn1, err := m.Method.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *AuditData_UpdateService) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.UpdateService != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuditData(dAtA, i, uint64(m.UpdateService.Size()))
		n2, err := m.UpdateService.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *AuditData_CreateVersion) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.CreateVersion != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuditData(dAtA, i, uint64(m.CreateVersion.Size()))
		n3, err := m.CreateVersion.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *UpdateServiceMethod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateServiceMethod) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Request != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuditData(dAtA, i, uint64(m.Request.Size()))
		n4, err := m.Request.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *CreateVersionMethod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateVersionMethod) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Request != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuditData(dAtA, i, uint64(m.Request.Size()))
		n5, err := m.Request.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func encodeFixed64AuditData(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32AuditData(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAuditData(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AuditData) Size() (n int) {
	var l int
	_ = l
	if m.Method != nil {
		n += m.Method.Size()
	}
	return n
}

func (m *AuditData_UpdateService) Size() (n int) {
	var l int
	_ = l
	if m.UpdateService != nil {
		l = m.UpdateService.Size()
		n += 1 + l + sovAuditData(uint64(l))
	}
	return n
}
func (m *AuditData_CreateVersion) Size() (n int) {
	var l int
	_ = l
	if m.CreateVersion != nil {
		l = m.CreateVersion.Size()
		n += 1 + l + sovAuditData(uint64(l))
	}
	return n
}
func (m *UpdateServiceMethod) Size() (n int) {
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovAuditData(uint64(l))
	}
	return n
}

func (m *CreateVersionMethod) Size() (n int) {
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovAuditData(uint64(l))
	}
	return n
}

func sovAuditData(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuditData(x uint64) (n int) {
	return sovAuditData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuditData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuditData
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
			return fmt.Errorf("proto: AuditData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuditData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditData
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
				return ErrInvalidLengthAuditData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &UpdateServiceMethod{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Method = &AuditData_UpdateService{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditData
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
				return ErrInvalidLengthAuditData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &CreateVersionMethod{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Method = &AuditData_CreateVersion{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuditData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuditData
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
func (m *UpdateServiceMethod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuditData
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
			return fmt.Errorf("proto: UpdateServiceMethod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateServiceMethod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditData
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
				return ErrInvalidLengthAuditData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &UpdateServiceRequest{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuditData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuditData
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
func (m *CreateVersionMethod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuditData
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
			return fmt.Errorf("proto: CreateVersionMethod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateVersionMethod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuditData
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
				return ErrInvalidLengthAuditData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &CreateVersionRequest{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuditData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuditData
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
func skipAuditData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuditData
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
					return 0, ErrIntOverflowAuditData
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
					return 0, ErrIntOverflowAuditData
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
				return 0, ErrInvalidLengthAuditData
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuditData
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
				next, err := skipAuditData(dAtA[start:])
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
	ErrInvalidLengthAuditData = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuditData   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("google/appengine/v1/audit_data.proto", fileDescriptorAuditData) }

var fileDescriptorAuditData = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xaf, 0xef, 0x50, 0xc0, 0x88, 0x0e, 0xe9, 0x40, 0xd5, 0x21, 0x42, 0x85, 0xa1, 0x2c,
	0x8e, 0x0a, 0x23, 0x2c, 0xa4, 0x0c, 0x2c, 0x48, 0x25, 0x08, 0x86, 0x2e, 0x91, 0x49, 0x8e, 0x8c,
	0xa5, 0x24, 0x36, 0x8e, 0x13, 0x89, 0x67, 0xe3, 0x05, 0x18, 0x79, 0x04, 0x94, 0x27, 0x41, 0xb6,
	0x43, 0x68, 0x51, 0x04, 0x8c, 0x39, 0xe7, 0xfb, 0xbf, 0xfc, 0xd2, 0x31, 0x3e, 0x62, 0x42, 0xb0,
	0x0c, 0x02, 0x2a, 0x25, 0x14, 0x8c, 0x17, 0x10, 0xd4, 0xf3, 0x80, 0x56, 0x29, 0xd7, 0x71, 0x4a,
	0x35, 0x25, 0x52, 0x09, 0x2d, 0xbc, 0x91, 0xa3, 0x48, 0x47, 0x91, 0x7a, 0x3e, 0x39, 0xec, 0x8d,
	0x76, 0x84, 0x4d, 0x4e, 0xfc, 0x16, 0xe2, 0x34, 0x37, 0x6b, 0x4e, 0xf3, 0x58, 0x8a, 0x8c, 0x27,
	0xcf, 0x6e, 0x3f, 0x7d, 0x41, 0x78, 0xe7, 0xc2, 0xfc, 0xee, 0x92, 0x6a, 0xea, 0xdd, 0xe0, 0x61,
	0x25, 0x53, 0xaa, 0x21, 0x2e, 0x41, 0xd5, 0x3c, 0x81, 0x31, 0x3a, 0x40, 0xb3, 0xdd, 0x93, 0x19,
	0xe9, 0x29, 0x40, 0xee, 0x2c, 0x7a, 0xeb, 0xc8, 0x6b, 0xd0, 0x8f, 0x22, 0xbd, 0xfa, 0x17, 0xed,
	0x55, 0xeb, 0x63, 0xa3, 0x4c, 0x14, 0x18, 0x65, 0x0d, 0xaa, 0xe4, 0xa2, 0x18, 0xff, 0xff, 0x41,
	0xb9, 0xb0, 0xe8, 0xbd, 0x23, 0xbf, 0x94, 0xc9, 0xfa, 0x38, 0xdc, 0xc6, 0x83, 0xdc, 0xae, 0xa6,
	0x2b, 0x3c, 0xea, 0x29, 0xe1, 0x2d, 0xf0, 0x96, 0x82, 0xa7, 0x0a, 0x4a, 0xdd, 0xf6, 0x3f, 0xfe,
	0xbd, 0x7f, 0xe4, 0x02, 0xd1, 0x67, 0xd2, 0xb8, 0x7b, 0xda, 0xfc, 0xd5, 0xbd, 0x11, 0xfd, 0xee,
	0x0e, 0xab, 0xd7, 0xc6, 0x47, 0x6f, 0x8d, 0x8f, 0xde, 0x1b, 0x1f, 0xe1, 0xfd, 0x44, 0xe4, 0x7d,
	0x92, 0x70, 0xd8, 0x5d, 0x66, 0x69, 0x8e, 0xb5, 0x44, 0xab, 0xf3, 0x16, 0x63, 0x22, 0xa3, 0x05,
	0x23, 0x42, 0xb1, 0x80, 0x41, 0x61, 0x4f, 0x19, 0xb8, 0x15, 0x95, 0xbc, 0xdc, 0x78, 0x12, 0x67,
	0xdd, 0xc7, 0xc3, 0xc0, 0x82, 0xa7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0xdf, 0xe8, 0x62,
	0x75, 0x02, 0x00, 0x00,
}
