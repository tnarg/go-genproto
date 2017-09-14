// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/monitored_resource.proto

/*
	Package monitoredres is a generated protocol buffer package.

	It is generated from these files:
		google/api/monitored_resource.proto

	It has these top-level messages:
		MonitoredResourceDescriptor
		MonitoredResource
*/
package monitoredres

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_api "google.golang.org/genproto/googleapis/api/label"

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

// An object that describes the schema of a [MonitoredResource][google.api.MonitoredResource] object using a
// type name and a set of labels.  For example, the monitored resource
// descriptor for Google Compute Engine VM instances has a type of
// `"gce_instance"` and specifies the use of the labels `"instance_id"` and
// `"zone"` to identify particular VM instances.
//
// Different APIs can support different monitored resource types. APIs generally
// provide a `list` method that returns the monitored resource descriptors used
// by the API.
type MonitoredResourceDescriptor struct {
	// Optional. The resource name of the monitored resource descriptor:
	// `"projects/{project_id}/monitoredResourceDescriptors/{type}"` where
	// {type} is the value of the `type` field in this object and
	// {project_id} is a project ID that provides API-specific context for
	// accessing the type.  APIs that do not use project information can use the
	// resource name format `"monitoredResourceDescriptors/{type}"`.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The monitored resource type. For example, the type
	// `"cloudsql_database"` represents databases in Google Cloud SQL.
	// The maximum length of this value is 256 characters.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Optional. A concise name for the monitored resource type that might be
	// displayed in user interfaces. It should be a Title Cased Noun Phrase,
	// without any article or other determiners. For example,
	// `"Google Cloud SQL Database"`.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. A detailed description of the monitored resource type that might
	// be used in documentation.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Required. A set of labels used to describe instances of this monitored
	// resource type. For example, an individual Google Cloud SQL database is
	// identified by values for the labels `"database_id"` and `"zone"`.
	Labels []*google_api.LabelDescriptor `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty"`
}

func (m *MonitoredResourceDescriptor) Reset()         { *m = MonitoredResourceDescriptor{} }
func (m *MonitoredResourceDescriptor) String() string { return proto.CompactTextString(m) }
func (*MonitoredResourceDescriptor) ProtoMessage()    {}
func (*MonitoredResourceDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptorMonitoredResource, []int{0}
}

func (m *MonitoredResourceDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetLabels() []*google_api.LabelDescriptor {
	if m != nil {
		return m.Labels
	}
	return nil
}

// An object representing a resource that can be used for monitoring, logging,
// billing, or other purposes. Examples include virtual machine instances,
// databases, and storage devices such as disks. The `type` field identifies a
// [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor] object that describes the resource's
// schema. Information in the `labels` field identifies the actual resource and
// its attributes according to the schema. For example, a particular Compute
// Engine VM instance could be represented by the following object, because the
// [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor] for `"gce_instance"` has labels
// `"instance_id"` and `"zone"`:
//
//     { "type": "gce_instance",
//       "labels": { "instance_id": "12345678901234",
//                   "zone": "us-central1-a" }}
type MonitoredResource struct {
	// Required. The monitored resource type. This field must match
	// the `type` field of a [MonitoredResourceDescriptor][google.api.MonitoredResourceDescriptor] object. For
	// example, the type of a Cloud SQL database is `"cloudsql_database"`.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Required. Values for all of the labels listed in the associated monitored
	// resource descriptor. For example, Cloud SQL databases use the labels
	// `"database_id"` and `"zone"`.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MonitoredResource) Reset()         { *m = MonitoredResource{} }
func (m *MonitoredResource) String() string { return proto.CompactTextString(m) }
func (*MonitoredResource) ProtoMessage()    {}
func (*MonitoredResource) Descriptor() ([]byte, []int) {
	return fileDescriptorMonitoredResource, []int{1}
}

func (m *MonitoredResource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MonitoredResource) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*MonitoredResourceDescriptor)(nil), "google.api.MonitoredResourceDescriptor")
	proto.RegisterType((*MonitoredResource)(nil), "google.api.MonitoredResource")
}
func (m *MonitoredResourceDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MonitoredResourceDescriptor) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMonitoredResource(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.DisplayName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMonitoredResource(dAtA, i, uint64(len(m.DisplayName)))
		i += copy(dAtA[i:], m.DisplayName)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMonitoredResource(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Labels) > 0 {
		for _, msg := range m.Labels {
			dAtA[i] = 0x22
			i++
			i = encodeVarintMonitoredResource(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMonitoredResource(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *MonitoredResource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MonitoredResource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMonitoredResource(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x12
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovMonitoredResource(uint64(len(k))) + 1 + len(v) + sovMonitoredResource(uint64(len(v)))
			i = encodeVarintMonitoredResource(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintMonitoredResource(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintMonitoredResource(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeFixed64MonitoredResource(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32MonitoredResource(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMonitoredResource(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MonitoredResourceDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMonitoredResource(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovMonitoredResource(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMonitoredResource(uint64(l))
	}
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovMonitoredResource(uint64(l))
		}
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMonitoredResource(uint64(l))
	}
	return n
}

func (m *MonitoredResource) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMonitoredResource(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovMonitoredResource(uint64(len(k))) + 1 + len(v) + sovMonitoredResource(uint64(len(v)))
			n += mapEntrySize + 1 + sovMonitoredResource(uint64(mapEntrySize))
		}
	}
	return n
}

func sovMonitoredResource(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMonitoredResource(x uint64) (n int) {
	return sovMonitoredResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MonitoredResourceDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoredResource
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
			return fmt.Errorf("proto: MonitoredResourceDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonitoredResourceDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
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
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
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
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
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
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
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
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, &google_api.LabelDescriptor{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
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
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoredResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMonitoredResource
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
func (m *MonitoredResource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoredResource
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
			return fmt.Errorf("proto: MonitoredResource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonitoredResource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
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
				return ErrInvalidLengthMonitoredResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredResource
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
				return ErrInvalidLengthMonitoredResource
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
						return ErrIntOverflowMonitoredResource
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
							return ErrIntOverflowMonitoredResource
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
						return ErrInvalidLengthMonitoredResource
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
							return ErrIntOverflowMonitoredResource
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
						return ErrInvalidLengthMonitoredResource
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMonitoredResource(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthMonitoredResource
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
			skippy, err := skipMonitoredResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMonitoredResource
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
func skipMonitoredResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMonitoredResource
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
					return 0, ErrIntOverflowMonitoredResource
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
					return 0, ErrIntOverflowMonitoredResource
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
				return 0, ErrInvalidLengthMonitoredResource
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMonitoredResource
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
				next, err := skipMonitoredResource(dAtA[start:])
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
	ErrInvalidLengthMonitoredResource = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMonitoredResource   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("google/api/monitored_resource.proto", fileDescriptorMonitoredResource)
}

var fileDescriptorMonitoredResource = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcb, 0x4a, 0xf3, 0x40,
	0x14, 0xc7, 0x99, 0xf4, 0x02, 0xdf, 0xe4, 0x43, 0x74, 0x90, 0x12, 0x5a, 0x28, 0xb5, 0x6e, 0xea,
	0x26, 0x01, 0xbb, 0xf1, 0xb2, 0x6a, 0x55, 0x44, 0x50, 0x29, 0x5d, 0xba, 0x29, 0xd3, 0xf6, 0x10,
	0x06, 0x93, 0x9c, 0x61, 0x26, 0x15, 0xf2, 0x00, 0xbe, 0x88, 0xe0, 0x73, 0xb8, 0x75, 0xe9, 0x23,
	0x48, 0x9f, 0xc2, 0xa5, 0x64, 0x32, 0xb5, 0x91, 0xb8, 0x3b, 0x97, 0xff, 0x39, 0xff, 0x5f, 0x72,
	0x86, 0x1e, 0x86, 0x88, 0x61, 0x04, 0x01, 0x97, 0x22, 0x88, 0x31, 0x11, 0x29, 0x2a, 0x58, 0xce,
	0x14, 0x68, 0x5c, 0xa9, 0x05, 0xf8, 0x52, 0x61, 0x8a, 0x8c, 0x16, 0x22, 0x9f, 0x4b, 0xd1, 0x6e,
	0x95, 0x06, 0x22, 0x3e, 0x87, 0xa8, 0xd0, 0xf4, 0xdf, 0x08, 0xed, 0xdc, 0x6d, 0x16, 0x4c, 0xed,
	0xfc, 0x25, 0xe8, 0x85, 0x12, 0x32, 0x45, 0xc5, 0x18, 0xad, 0xa7, 0x99, 0x04, 0x8f, 0xf4, 0xc8,
	0xe0, 0xdf, 0xd4, 0xc4, 0xec, 0x80, 0xfe, 0x5f, 0x0a, 0x2d, 0x23, 0x9e, 0xcd, 0x12, 0x1e, 0x83,
	0xe7, 0x98, 0x9e, 0x6b, 0x6b, 0xf7, 0x3c, 0x06, 0xd6, 0xa3, 0xee, 0xd2, 0x2e, 0x11, 0x98, 0x78,
	0x35, 0xab, 0xd8, 0x96, 0xd8, 0x90, 0x36, 0x0d, 0x87, 0xf6, 0xea, 0xbd, 0xda, 0xc0, 0x3d, 0xee,
	0xf8, 0x5b, 0x5a, 0xff, 0x36, 0xef, 0x6c, 0x29, 0xa6, 0x56, 0x9a, 0xd3, 0x18, 0xc7, 0x46, 0x41,
	0x93, 0xc7, 0xfd, 0x57, 0x42, 0xf7, 0x2a, 0x5f, 0xf0, 0x27, 0xf7, 0xe8, 0xc7, 0xd2, 0x31, 0x96,
	0x47, 0x65, 0xcb, 0xca, 0x8a, 0x02, 0x42, 0x5f, 0x25, 0xa9, 0xca, 0x36, 0x00, 0xed, 0x53, 0xea,
	0x96, 0xca, 0x6c, 0x97, 0xd6, 0x1e, 0x21, 0xb3, 0x26, 0x79, 0xc8, 0xf6, 0x69, 0xe3, 0x89, 0x47,
	0xab, 0xcd, 0x4f, 0x29, 0x92, 0x33, 0xe7, 0x84, 0x8c, 0x9f, 0xc9, 0xfb, 0xba, 0x4b, 0x3e, 0xd6,
	0x5d, 0xf2, 0xb9, 0xee, 0x12, 0xba, 0xb3, 0xc0, 0xb8, 0xe4, 0x3f, 0x6e, 0x55, 0x00, 0x26, 0xf9,
	0x81, 0x26, 0xe4, 0xe1, 0xc2, 0xaa, 0x42, 0x8c, 0x78, 0x12, 0xfa, 0xa8, 0xc2, 0x20, 0x84, 0xc4,
	0x9c, 0x2f, 0x28, 0x5a, 0x5c, 0x0a, 0xfd, 0xfb, 0x29, 0x28, 0xd0, 0xe7, 0xe5, 0xe4, 0x8b, 0x90,
	0x17, 0xa7, 0x7e, 0x3d, 0x9a, 0xdc, 0xcc, 0x9b, 0x66, 0x72, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0x40, 0xf5, 0xa1, 0x4f, 0x43, 0x02, 0x00, 0x00,
}
