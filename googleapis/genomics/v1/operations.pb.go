// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/genomics/v1/operations.proto

package genomics

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf5 "github.com/gogo/protobuf/types"
import google_protobuf6 "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Metadata describing an [Operation][google.longrunning.Operation].
type OperationMetadata struct {
	// The Google Cloud Project in which the job is scoped.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The time at which the job was submitted to the Genomics service.
	CreateTime *google_protobuf6.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// The time at which the job began to run.
	StartTime *google_protobuf6.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// The time at which the job stopped running.
	EndTime *google_protobuf6.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// The original request that started the operation. Note that this will be in
	// current version of the API. If the operation was started with v1beta2 API
	// and a GetOperation is performed on v1 API, a v1 request will be returned.
	Request *google_protobuf5.Any `protobuf:"bytes,5,opt,name=request" json:"request,omitempty"`
	// Optional event messages that were generated during the job's execution.
	// This also contains any warnings that were generated during import
	// or export.
	Events []*OperationEvent `protobuf:"bytes,6,rep,name=events" json:"events,omitempty"`
	// This field is deprecated. Use `labels` instead. Optionally provided by the
	// caller when submitting the request that creates the operation.
	ClientId string `protobuf:"bytes,7,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Runtime metadata on this Operation.
	RuntimeMetadata *google_protobuf5.Any `protobuf:"bytes,8,opt,name=runtime_metadata,json=runtimeMetadata" json:"runtime_metadata,omitempty"`
	// Optionally provided by the caller when submitting the request that creates
	// the operation.
	Labels map[string]string `protobuf:"bytes,9,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *OperationMetadata) Reset()                    { *m = OperationMetadata{} }
func (m *OperationMetadata) String() string            { return proto.CompactTextString(m) }
func (*OperationMetadata) ProtoMessage()               {}
func (*OperationMetadata) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{0} }

func (m *OperationMetadata) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *OperationMetadata) GetCreateTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *OperationMetadata) GetStartTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OperationMetadata) GetEndTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *OperationMetadata) GetRequest() *google_protobuf5.Any {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *OperationMetadata) GetEvents() []*OperationEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *OperationMetadata) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *OperationMetadata) GetRuntimeMetadata() *google_protobuf5.Any {
	if m != nil {
		return m.RuntimeMetadata
	}
	return nil
}

func (m *OperationMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// An event that occurred during an [Operation][google.longrunning.Operation].
type OperationEvent struct {
	// Optional time of when event started.
	StartTime *google_protobuf6.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// Optional time of when event finished. An event can have a start time and no
	// finish time. If an event has a finish time, there must be a start time.
	EndTime *google_protobuf6.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// Required description of event.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *OperationEvent) Reset()                    { *m = OperationEvent{} }
func (m *OperationEvent) String() string            { return proto.CompactTextString(m) }
func (*OperationEvent) ProtoMessage()               {}
func (*OperationEvent) Descriptor() ([]byte, []int) { return fileDescriptorOperations, []int{1} }

func (m *OperationEvent) GetStartTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OperationEvent) GetEndTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *OperationEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*OperationMetadata)(nil), "google.genomics.v1.OperationMetadata")
	proto.RegisterType((*OperationEvent)(nil), "google.genomics.v1.OperationEvent")
}
func (m *OperationMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OperationMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProjectId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOperations(dAtA, i, uint64(len(m.ProjectId)))
		i += copy(dAtA[i:], m.ProjectId)
	}
	if m.CreateTime != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOperations(dAtA, i, uint64(m.CreateTime.Size()))
		n1, err := m.CreateTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.StartTime != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOperations(dAtA, i, uint64(m.StartTime.Size()))
		n2, err := m.StartTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.EndTime != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOperations(dAtA, i, uint64(m.EndTime.Size()))
		n3, err := m.EndTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Request != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintOperations(dAtA, i, uint64(m.Request.Size()))
		n4, err := m.Request.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0x32
			i++
			i = encodeVarintOperations(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ClientId) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintOperations(dAtA, i, uint64(len(m.ClientId)))
		i += copy(dAtA[i:], m.ClientId)
	}
	if m.RuntimeMetadata != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintOperations(dAtA, i, uint64(m.RuntimeMetadata.Size()))
		n5, err := m.RuntimeMetadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x4a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovOperations(uint64(len(k))) + 1 + len(v) + sovOperations(uint64(len(v)))
			i = encodeVarintOperations(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintOperations(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintOperations(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *OperationEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OperationEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StartTime != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOperations(dAtA, i, uint64(m.StartTime.Size()))
		n6, err := m.StartTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.EndTime != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOperations(dAtA, i, uint64(m.EndTime.Size()))
		n7, err := m.EndTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOperations(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	return i, nil
}

func encodeFixed64Operations(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Operations(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintOperations(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *OperationMetadata) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProjectId)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.CreateTime != nil {
		l = m.CreateTime.Size()
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.StartTime != nil {
		l = m.StartTime.Size()
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.EndTime != nil {
		l = m.EndTime.Size()
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovOperations(uint64(l))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovOperations(uint64(l))
		}
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.RuntimeMetadata != nil {
		l = m.RuntimeMetadata.Size()
		n += 1 + l + sovOperations(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovOperations(uint64(len(k))) + 1 + len(v) + sovOperations(uint64(len(v)))
			n += mapEntrySize + 1 + sovOperations(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *OperationEvent) Size() (n int) {
	var l int
	_ = l
	if m.StartTime != nil {
		l = m.StartTime.Size()
		n += 1 + l + sovOperations(uint64(l))
	}
	if m.EndTime != nil {
		l = m.EndTime.Size()
		n += 1 + l + sovOperations(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovOperations(uint64(l))
	}
	return n
}

func sovOperations(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOperations(x uint64) (n int) {
	return sovOperations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OperationMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: OperationMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OperationMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreateTime == nil {
				m.CreateTime = &google_protobuf6.Timestamp{}
			}
			if err := m.CreateTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartTime == nil {
				m.StartTime = &google_protobuf6.Timestamp{}
			}
			if err := m.StartTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndTime == nil {
				m.EndTime = &google_protobuf6.Timestamp{}
			}
			if err := m.EndTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &google_protobuf5.Any{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &OperationEvent{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RuntimeMetadata == nil {
				m.RuntimeMetadata = &google_protobuf5.Any{}
			}
			if err := m.RuntimeMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
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
						return ErrIntOverflowOperations
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
							return ErrIntOverflowOperations
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
						return ErrInvalidLengthOperations
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
							return ErrIntOverflowOperations
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
						return ErrInvalidLengthOperations
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipOperations(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthOperations
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
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperations
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
func (m *OperationEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperations
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
			return fmt.Errorf("proto: OperationEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OperationEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartTime == nil {
				m.StartTime = &google_protobuf6.Timestamp{}
			}
			if err := m.StartTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndTime == nil {
				m.EndTime = &google_protobuf6.Timestamp{}
			}
			if err := m.EndTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperations
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
				return ErrInvalidLengthOperations
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOperations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperations
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
func skipOperations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOperations
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
					return 0, ErrIntOverflowOperations
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
					return 0, ErrIntOverflowOperations
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
				return 0, ErrInvalidLengthOperations
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOperations
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
				next, err := skipOperations(dAtA[start:])
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
	ErrInvalidLengthOperations = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOperations   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("google/genomics/v1/operations.proto", fileDescriptorOperations) }

var fileDescriptorOperations = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xe5, 0x76, 0x6b, 0x9b, 0x2f, 0x12, 0x1b, 0xd6, 0x84, 0x42, 0x81, 0x52, 0x95, 0x4b,
	0x4f, 0x8e, 0x3a, 0x84, 0xc4, 0xba, 0x03, 0x62, 0xd2, 0x0e, 0x95, 0x40, 0x4c, 0x11, 0x27, 0x2e,
	0x95, 0x9b, 0x7c, 0x44, 0x81, 0xc4, 0x0e, 0xb6, 0x1b, 0xa9, 0xef, 0xc3, 0x0b, 0xf0, 0x16, 0x1c,
	0x79, 0x04, 0xd4, 0xa7, 0xe0, 0x88, 0xe2, 0x38, 0x55, 0xd9, 0xd0, 0x86, 0xb8, 0xc5, 0xdf, 0xf7,
	0xff, 0xd9, 0x7f, 0x7f, 0xff, 0x18, 0x9e, 0xa5, 0x52, 0xa6, 0x39, 0x86, 0x29, 0x0a, 0x59, 0x64,
	0xb1, 0x0e, 0xab, 0x59, 0x28, 0x4b, 0x54, 0xdc, 0x64, 0x52, 0x68, 0x56, 0x2a, 0x69, 0x24, 0xa5,
	0x8d, 0x88, 0xb5, 0x22, 0x56, 0xcd, 0x86, 0x8f, 0x1d, 0xc8, 0xcb, 0x2c, 0xe4, 0x42, 0x48, 0xb3,
	0x4f, 0x0c, 0x1f, 0xba, 0xae, 0x5d, 0xad, 0xd6, 0x1f, 0x43, 0x2e, 0x36, 0xae, 0xf5, 0xf4, 0x7a,
	0xcb, 0x64, 0x05, 0x6a, 0xc3, 0x8b, 0xb2, 0x11, 0x4c, 0xbe, 0x1d, 0xc0, 0xfd, 0x77, 0xad, 0x85,
	0xb7, 0x68, 0x78, 0xc2, 0x0d, 0xa7, 0x4f, 0x00, 0x4a, 0x25, 0x3f, 0x61, 0x6c, 0x96, 0x59, 0x12,
	0x90, 0x31, 0x99, 0x7a, 0x91, 0xe7, 0x2a, 0x8b, 0x84, 0x9e, 0x83, 0x1f, 0x2b, 0xe4, 0x06, 0x97,
	0xf5, 0x76, 0x41, 0x67, 0x4c, 0xa6, 0xfe, 0xe9, 0x90, 0x39, 0xe3, 0xed, 0x59, 0xec, 0x7d, 0x7b,
	0x56, 0x04, 0x8d, 0xbc, 0x2e, 0xd0, 0x33, 0x00, 0x6d, 0xb8, 0x32, 0x0d, 0xdb, 0xbd, 0x93, 0xf5,
	0xac, 0xda, 0xa2, 0x2f, 0x60, 0x80, 0x22, 0x69, 0xc0, 0x83, 0x3b, 0xc1, 0x3e, 0x8a, 0xc4, 0x62,
	0x0c, 0xfa, 0x0a, 0xbf, 0xac, 0x51, 0x9b, 0xe0, 0xd0, 0x52, 0x27, 0x37, 0xa8, 0xd7, 0x62, 0x13,
	0xb5, 0x22, 0x3a, 0x87, 0x1e, 0x56, 0x28, 0x8c, 0x0e, 0x7a, 0xe3, 0xee, 0xd4, 0x3f, 0x9d, 0xb0,
	0x9b, 0x91, 0xb0, 0xdd, 0xd0, 0x2e, 0x6b, 0x69, 0xe4, 0x08, 0xfa, 0x08, 0xbc, 0x38, 0xcf, 0x50,
	0xd8, 0xc1, 0xf5, 0xed, 0xe0, 0x06, 0x4d, 0x61, 0x91, 0xd0, 0x57, 0x70, 0xac, 0xd6, 0xa2, 0xb6,
	0xbf, 0x2c, 0xdc, 0xa8, 0x83, 0xc1, 0x2d, 0x8e, 0x8e, 0x9c, 0x7a, 0x97, 0xcb, 0x02, 0x7a, 0x39,
	0x5f, 0x61, 0xae, 0x03, 0xcf, 0x3a, 0x9b, 0xdd, 0xea, 0xac, 0xc5, 0xd8, 0x1b, 0xcb, 0x5c, 0x0a,
	0xa3, 0x36, 0x91, 0xdb, 0x60, 0x78, 0x06, 0xfe, 0x5e, 0x99, 0x1e, 0x43, 0xf7, 0x33, 0x6e, 0x5c,
	0xd4, 0xf5, 0x27, 0x3d, 0x81, 0xc3, 0x8a, 0xe7, 0xeb, 0x26, 0x5e, 0x2f, 0x6a, 0x16, 0xf3, 0xce,
	0x4b, 0x32, 0xf9, 0x4a, 0xe0, 0xde, 0x9f, 0xd7, 0xbf, 0x16, 0x2a, 0xf9, 0xdf, 0x50, 0x3b, 0xff,
	0x1e, 0xea, 0x18, 0xfc, 0x04, 0x75, 0xac, 0xb2, 0xb2, 0x76, 0x61, 0xff, 0x23, 0x2f, 0xda, 0x2f,
	0x5d, 0x54, 0xdf, 0xb7, 0x23, 0xf2, 0x63, 0x3b, 0x22, 0x3f, 0xb7, 0x23, 0x02, 0x0f, 0x62, 0x59,
	0xfc, 0x65, 0x5a, 0x17, 0x47, 0xbb, 0x9b, 0xe8, 0xab, 0xfa, 0xb8, 0x2b, 0xf2, 0x61, 0xde, 0xca,
	0x64, 0xce, 0x45, 0xca, 0xa4, 0x4a, 0xeb, 0x17, 0x6b, 0xcd, 0x84, 0x4d, 0x8b, 0x97, 0x99, 0xde,
	0x7f, 0xc5, 0xe7, 0xed, 0xf7, 0x2f, 0x42, 0x56, 0x3d, 0xab, 0x7c, 0xfe, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x5e, 0xba, 0x38, 0x60, 0xee, 0x03, 0x00, 0x00,
}
